package urtransform

import (
	"fmt"
	"net/url"
	"strings"
)

// URLNodeType is of URLNode type
type URLNodeType int

const (
	Scheme URLNodeType = iota
	Location
	Path
	User
	Param
	ParamValue
	String
)

// URLNode - URLTransformRule node
type URLNode struct {
	Type URLNodeType
	Name string
}

// URL Transform rule
type URLTransformRule []URLNode

// NewURLTransformRule create transform rule. For example original URL
// {scheme}:://{location}/{path}?{param_name.name1}={param_value.name1}&{param_name.name2}={param_value.name2}
func NewURLTransformRule(urltransform string) (URLTransformRule, error) {
	var ut URLTransformRule
	if len(urltransform) > 0 {
		ut = make(URLTransformRule, 0, 10)
		s := urltransform[0:]
		for {
			start := strings.IndexByte(s, '{')
			end := strings.IndexByte(s, '}')
			if start > 0 {
				// String
				ut = append(ut, URLNode{String, s[0:start]})
			} else if start == -1 {
				if end == -1 {
					// String at the end
					ut = append(ut, URLNode{String, s})
					return ut, nil
				} else {
					return nil, fmt.Errorf("unclosed { in in rule node: %s", s)
				}
			}
			if end == -1 {
				return nil, fmt.Errorf("unclosed { in URL transform rule")
			}
			end++
			node := s[start:end]
			switch node {
			case "{scheme}":
				ut = append(ut, URLNode{Scheme, ""})
			case "{location}":
				ut = append(ut, URLNode{Location, ""})
			case "{user}":
				ut = append(ut, URLNode{User, ""})
			case "{path}":
				ut = append(ut, URLNode{Path, ""})
			default:
				if strings.HasPrefix(node, "{param.") {
					ut = append(ut, URLNode{Param, node[7 : len(node)-1]})
				} else if strings.HasPrefix(node, "{param_value.") {
					ut = append(ut, URLNode{ParamValue, node[13 : len(node)-1]})
				} else {
					return nil, fmt.Errorf("unknown node: %s", node)
				}
			}
			if end == len(s) {
				break
			}
			s = s[end:]
		}
	}
	return ut, nil
}

// URLTransformRuleEqual compare URLTransformRules
func URLTransformRuleEqual(a, b URLTransformRule) bool {
	if &a == &b {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Type != b[i].Type || a[i].Name != b[i].Name {
			return false
		}
	}
	return true
}

// URLTransform url with rule with optional decode
func URLTransform(surl string, tr URLTransformRule) (string, error) {
	u, err := url.ParseRequestURI(surl)
	if err != nil {
		return surl, err
	}
	var b strings.Builder
	for i := range tr {
		switch tr[i].Type {
		case Scheme:
			b.WriteString(u.Scheme)
		case User:
			if u.User == nil {
				b.WriteString("anonymous")
			} else {
				b.WriteString(u.User.Username())
			}
		case Location:
			b.WriteString(u.Hostname())
			p := u.Port()
			if p != "" {
				b.WriteString(":")
				b.WriteString(p)
			}
		case Path:
			b.WriteString(u.Path)
		case String:
			b.WriteString(tr[i].Name)
		case Param:
			params, found := u.Query()[tr[i].Name]
			if found {
				for k := range params {
					if k > 0 {
						b.WriteString("&")
					}
					b.WriteString(tr[i].Name)
					b.WriteString("=")
					b.WriteString(params[k])
				}
			} else {
				return surl, fmt.Errorf("param not found: %s", tr[i].Name)
			}
		case ParamValue:
			params, found := u.Query()[tr[i].Name]
			if found {
				for k := range params {
					if k > 0 {
						b.WriteString("&")
					}
					b.WriteString(params[k])
				}
			} else {
				return surl, fmt.Errorf("param not found: %s", tr[i].Name)
			}
		default:
			return surl, fmt.Errorf("unhandled node type: %d", tr[i].Type)
		}
	}
	return b.String(), nil
}

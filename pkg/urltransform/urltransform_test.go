package urtransform

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestCorrectRule(t *testing.T) {
	var tests = []struct {
		n        string           // input
		expected URLTransformRule // expected result
	}{
		{
			"{scheme}:://{user}@{location}/{path}?{param.name1}&{param_value.name2}",
			URLTransformRule{
				{Scheme, ""},
				{String, ":://"},
				{User, ""},
				{String, "@"},
				{Location, ""},
				{String, "/"},
				{Path, ""},
				{String, "?"},
				{Param, "name1"},
				{String, "&"},
				{ParamValue, "name2"},
			},
		},
	}

	for _, tt := range tests {
		actual, err := NewURLTransformRule(tt.n)
		if err != nil {
			t.Errorf("NewURLTransformRule(%v): error %s", tt.n, err.Error())
		}
		if !URLTransformRuleEqual(tt.expected, actual) {
			t.Errorf("NewURLTransformRule(%v): expected %v, actual %v", tt.n, tt.expected, actual)
		}
	}
}

func TestInvalidRule(t *testing.T) {
	var tests = []struct {
		input string
		error string
	}{
		{"{scheme}:://location}", "unclosed { in in rule node: :://location}"},
		{"{scheme}:://{location/{path}", "unknown node: {location/{path}"},
	}

	for _, tt := range tests {
		actual, err := NewURLTransformRule(tt.input)
		if err == nil {
			t.Errorf("NewURLTransformRule(%s): mailformed rule, got %v", tt.input, actual)
		} else if tt.error != err.Error() {
			t.Errorf("NewURLTransformRule(%s): %s", tt.input, err.Error())
		}
	}
}

func TestURLTranformRule(t *testing.T) {
	str := "{scheme}://{user}@{location}{path}?{param.name1}&{param_value.name2}"
	tr, err := NewURLTransformRule(str)
	if err != nil {
		t.Errorf("NewURLTransformRule(%v): error %s", str, err.Error())
	}
	var tests = []struct {
		input  string
		output string
	}{
		{
			"http://user@127.0.0.1:8080/render?name1=test1&name1=test2&name2=test3&name2=test4&name3=test5",
			"http://user@127.0.0.1:8080/render?name1=test1&name1=test2&test3&test4",
		},
	}

	for _, tt := range tests {
		actual, err := URLTransform(tt.input, tr)
		if err != nil {
			t.Errorf("URLTransform(%s, ..): %s", tt.input, err.Error())
		} else {
			assert.Equal(t, tt.output, actual, "incorrect URLTransform result")
		}
	}
}

func BenchmarkURLTRansformRule(b *testing.B) {
	str := "{scheme}://{user}@{location}{path}?{param.name1}&{param_value.name2}"
	tr, err := NewURLTransformRule(str)
	if err != nil {
		b.Errorf("NewURLTransformRule(%v): error %s", str, err.Error())
	}
	input := "http://user@127.0.0.1:8080/render?name1=test1&name1=test2&name2=test3&name2=test4&name3=test5"
	output := "http://user@127.0.0.1:8080/render?name1=test1&name1=test2&test3&test4"
	for i := 0; i < b.N; i++ {
		actual, err := URLTransform(input, tr)
		if err != nil {
			b.Errorf("URLTransform(%s, ..): %s", input, err.Error())
		} else {
			assert.Equal(b, output, actual, "incorrect URLTransform result")
		}
	}
}

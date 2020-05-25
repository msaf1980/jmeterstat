// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package aggstat

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat(in *jlexer.Lexer, out *URLAggStat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Stat":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Stat = make(map[string]*AggStat)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 *AggStat
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(AggStat)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					(out.Stat)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "SumStat":
			(out.SumStat).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat(out *jwriter.Writer, in URLAggStat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Stat\":"
		out.RawString(prefix[1:])
		if in.Stat == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Stat {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if v2Value == nil {
					out.RawString("null")
				} else {
					(*v2Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"SumStat\":"
		out.RawString(prefix)
		(in.SumStat).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v URLAggStat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v URLAggStat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *URLAggStat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *URLAggStat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat(l, v)
}
func easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat1(in *jlexer.Lexer, out *LabelURLAggStat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Started":
			out.Started = int64(in.Int64())
		case "Ended":
			out.Ended = int64(in.Int64())
		case "Name":
			out.Name = string(in.String())
		case "Stat":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Stat = make(map[string]*URLAggStat)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 *URLAggStat
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(URLAggStat)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					(out.Stat)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat1(out *jwriter.Writer, in LabelURLAggStat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Started\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Started))
	}
	{
		const prefix string = ",\"Ended\":"
		out.RawString(prefix)
		out.Int64(int64(in.Ended))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Stat\":"
		out.RawString(prefix)
		if in.Stat == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v4First := true
			for v4Name, v4Value := range in.Stat {
				if v4First {
					v4First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v4Name))
				out.RawByte(':')
				if v4Value == nil {
					out.RawString("null")
				} else {
					(*v4Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LabelURLAggStat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LabelURLAggStat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LabelURLAggStat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LabelURLAggStat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat1(l, v)
}
func easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat2(in *jlexer.Lexer, out *AggStatNode) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Min":
			out.Min = float64(in.Float64())
		case "Max":
			out.Max = float64(in.Float64())
		case "Mean":
			out.Mean = float64(in.Float64())
		case "P90":
			out.P90 = float64(in.Float64())
		case "P95":
			out.P95 = float64(in.Float64())
		case "P99":
			out.P99 = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat2(out *jwriter.Writer, in AggStatNode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Min\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Min))
	}
	{
		const prefix string = ",\"Max\":"
		out.RawString(prefix)
		out.Float64(float64(in.Max))
	}
	{
		const prefix string = ",\"Mean\":"
		out.RawString(prefix)
		out.Float64(float64(in.Mean))
	}
	{
		const prefix string = ",\"P90\":"
		out.RawString(prefix)
		out.Float64(float64(in.P90))
	}
	{
		const prefix string = ",\"P95\":"
		out.RawString(prefix)
		out.Float64(float64(in.P95))
	}
	{
		const prefix string = ",\"P99\":"
		out.RawString(prefix)
		out.Float64(float64(in.P99))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AggStatNode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AggStatNode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AggStatNode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AggStatNode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat2(l, v)
}
func easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat3(in *jlexer.Lexer, out *AggStat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Started":
			out.Started = int64(in.Int64())
		case "Ended":
			out.Ended = int64(in.Int64())
		case "Count":
			out.Count = uint64(in.Uint64())
		case "Elapsed":
			(out.Elapsed).UnmarshalEasyJSON(in)
		case "Connect":
			(out.Connect).UnmarshalEasyJSON(in)
		case "Bytes":
			(out.Bytes).UnmarshalEasyJSON(in)
		case "SentBytes":
			(out.SentBytes).UnmarshalEasyJSON(in)
		case "Success":
			out.Success = uint64(in.Uint64())
		case "SuccessCodes":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.SuccessCodes = make(map[string]uint64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v5 uint64
					v5 = uint64(in.Uint64())
					(out.SuccessCodes)[key] = v5
					in.WantComma()
				}
				in.Delim('}')
			}
		case "ErrorCodes":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.ErrorCodes = make(map[string]uint64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v6 uint64
					v6 = uint64(in.Uint64())
					(out.ErrorCodes)[key] = v6
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat3(out *jwriter.Writer, in AggStat) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Started\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Started))
	}
	{
		const prefix string = ",\"Ended\":"
		out.RawString(prefix)
		out.Int64(int64(in.Ended))
	}
	{
		const prefix string = ",\"Count\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Count))
	}
	{
		const prefix string = ",\"Elapsed\":"
		out.RawString(prefix)
		(in.Elapsed).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Connect\":"
		out.RawString(prefix)
		(in.Connect).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Bytes\":"
		out.RawString(prefix)
		(in.Bytes).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"SentBytes\":"
		out.RawString(prefix)
		(in.SentBytes).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"Success\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Success))
	}
	{
		const prefix string = ",\"SuccessCodes\":"
		out.RawString(prefix)
		if in.SuccessCodes == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v7First := true
			for v7Name, v7Value := range in.SuccessCodes {
				if v7First {
					v7First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v7Name))
				out.RawByte(':')
				out.Uint64(uint64(v7Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"ErrorCodes\":"
		out.RawString(prefix)
		if in.ErrorCodes == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.ErrorCodes {
				if v8First {
					v8First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v8Name))
				out.RawByte(':')
				out.Uint64(uint64(v8Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AggStat) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AggStat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF1aca853EncodeGithubComMsaf1980JmeterstatPkgAggstat3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AggStat) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AggStat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF1aca853DecodeGithubComMsaf1980JmeterstatPkgAggstat3(l, v)
}

// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package benchmark

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

func easyjson21677a1cDecodeGithubComSapphiRedJsonConstantiaterBenchmarks(in *jlexer.Lexer, out *SmallPayload) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "St":
			out.St = int(in.Int())
		case "Sid":
			out.Sid = int(in.Int())
		case "Tt":
			out.Tt = string(in.String())
		case "Gr":
			out.Gr = int(in.Int())
		case "Uuid":
			out.Uuid = string(in.String())
		case "Ip":
			out.Ip = string(in.String())
		case "Ua":
			out.Ua = string(in.String())
		case "Tz":
			out.Tz = int(in.Int())
		case "V":
			out.V = int(in.Int())
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
func easyjson21677a1cEncodeGithubComSapphiRedJsonConstantiaterBenchmarks(out *jwriter.Writer, in SmallPayload) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"St\":"
		out.RawString(prefix[1:])
		out.Int(int(in.St))
	}
	{
		const prefix string = ",\"Sid\":"
		out.RawString(prefix)
		out.Int(int(in.Sid))
	}
	{
		const prefix string = ",\"Tt\":"
		out.RawString(prefix)
		out.String(string(in.Tt))
	}
	{
		const prefix string = ",\"Gr\":"
		out.RawString(prefix)
		out.Int(int(in.Gr))
	}
	{
		const prefix string = ",\"Uuid\":"
		out.RawString(prefix)
		out.String(string(in.Uuid))
	}
	{
		const prefix string = ",\"Ip\":"
		out.RawString(prefix)
		out.String(string(in.Ip))
	}
	{
		const prefix string = ",\"Ua\":"
		out.RawString(prefix)
		out.String(string(in.Ua))
	}
	{
		const prefix string = ",\"Tz\":"
		out.RawString(prefix)
		out.Int(int(in.Tz))
	}
	{
		const prefix string = ",\"V\":"
		out.RawString(prefix)
		out.Int(int(in.V))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SmallPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson21677a1cEncodeGithubComSapphiRedJsonConstantiaterBenchmarks(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SmallPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson21677a1cDecodeGithubComSapphiRedJsonConstantiaterBenchmarks(l, v)
}

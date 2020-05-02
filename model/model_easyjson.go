// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjsonC80ae7adDecodeGithubComRomapres2010HttpserverModel(in *jlexer.Lexer, out *Emp) {
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
		case "empNo":
			out.Empno = int(in.Int())
		case "empName":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ename).UnmarshalJSON(data))
			}
		case "job":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Job).UnmarshalJSON(data))
			}
		case "mgr":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Mgr).UnmarshalJSON(data))
			}
		case "hiredate":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Hiredate).UnmarshalJSON(data))
			}
		case "sal":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Sal).UnmarshalJSON(data))
			}
		case "comm":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Comm).UnmarshalJSON(data))
			}
		case "deptNumber":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Deptno).UnmarshalJSON(data))
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
func easyjsonC80ae7adEncodeGithubComRomapres2010HttpserverModel(out *jwriter.Writer, in Emp) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"empNo\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Empno))
	}
	{
		const prefix string = ",\"empName\":"
		out.RawString(prefix)
		out.Raw((in.Ename).MarshalJSON())
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.Raw((in.Job).MarshalJSON())
	}
	if true {
		const prefix string = ",\"mgr\":"
		out.RawString(prefix)
		out.Raw((in.Mgr).MarshalJSON())
	}
	{
		const prefix string = ",\"hiredate\":"
		out.RawString(prefix)
		out.Raw((in.Hiredate).MarshalJSON())
	}
	{
		const prefix string = ",\"sal\":"
		out.RawString(prefix)
		out.Raw((in.Sal).MarshalJSON())
	}
	{
		const prefix string = ",\"comm\":"
		out.RawString(prefix)
		out.Raw((in.Comm).MarshalJSON())
	}
	{
		const prefix string = ",\"deptNumber\":"
		out.RawString(prefix)
		out.Raw((in.Deptno).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Emp) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComRomapres2010HttpserverModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Emp) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComRomapres2010HttpserverModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Emp) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComRomapres2010HttpserverModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Emp) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComRomapres2010HttpserverModel(l, v)
}
func easyjsonC80ae7adDecodeGithubComRomapres2010HttpserverModel1(in *jlexer.Lexer, out *Dept) {
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
		case "deptNumber":
			out.Deptno = int(in.Int())
		case "deptName":
			out.Dname = string(in.String())
		case "deptLocation":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Loc).UnmarshalJSON(data))
			}
		case "emps":
			if in.IsNull() {
				in.Skip()
				out.Emps = nil
			} else {
				in.Delim('[')
				if out.Emps == nil {
					if !in.IsDelim(']') {
						out.Emps = make([]*Emp, 0, 8)
					} else {
						out.Emps = []*Emp{}
					}
				} else {
					out.Emps = (out.Emps)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Emp
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Emp)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Emps = append(out.Emps, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonC80ae7adEncodeGithubComRomapres2010HttpserverModel1(out *jwriter.Writer, in Dept) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"deptNumber\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Deptno))
	}
	{
		const prefix string = ",\"deptName\":"
		out.RawString(prefix)
		out.String(string(in.Dname))
	}
	{
		const prefix string = ",\"deptLocation\":"
		out.RawString(prefix)
		out.Raw((in.Loc).MarshalJSON())
	}
	if len(in.Emps) != 0 {
		const prefix string = ",\"emps\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.Emps {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Dept) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComRomapres2010HttpserverModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Dept) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComRomapres2010HttpserverModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Dept) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComRomapres2010HttpserverModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Dept) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComRomapres2010HttpserverModel1(l, v)
}

// Code generated by json-constantiater DO NOT EDIT

package benchmark

import "github.com/sapphi-red/json-constantiater/lib"

func (t *SmallPayload) NewJsonMarshal() []byte {
	tmpPtr := lib.GetFromPool()
	tmp := *tmpPtr
	tmp = t.AppendJsonString(tmp)
	res := make([]byte, len(tmp))
	copy(res, tmp)
	*tmpPtr = tmp
	lib.PutToPool(tmpPtr)
	return res
}

//go:nosplit
func (t *SmallPayload) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"st\":"...)
	res = lib.AppendInt(res, t.St)
	res = append(res, ",\"sid\":"...)
	res = lib.AppendInt(res, t.Sid)
	res = append(res, ",\"tt\":\""...)
	res = lib.AppendByteWithEscape(res, t.Tt)
	res = append(res, "\",\"gr\":"...)
	res = lib.AppendInt(res, t.Gr)
	res = append(res, ",\"uuid\":\""...)
	res = append(res, t.Uuid...)
	res = append(res, "\",\"ip\":\""...)
	res = append(res, t.Ip...)
	res = append(res, "\",\"ua\":\""...)
	res = lib.AppendByteWithEscape(res, t.Ua)
	res = append(res, "\",\"tz\":"...)
	res = lib.AppendInt(res, t.Tz)
	res = append(res, ",\"v\":"...)
	res = lib.AppendInt(res, t.V)
	res = append(res, '}')
	return res
}

//go:nosplit
func (t *SmallPayload) JsonLen() int {
	l := 2
	l += 20
	l += 2 + 2 + 1 + 1
	l += 20
	l += 2 + 3 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Tt)
	l += 2 + 2 + 1 + 1
	l += 20
	l += 2 + 2 + 1 + 1
	l += 38
	l += 2 + 4 + 1 + 1
	l += 17
	l += 2 + 2 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Ua)
	l += 2 + 2 + 1 + 1
	l += 20
	l += 2 + 2 + 1 + 1
	l += 2
	l += 2 + 1 + 1 + 1
	return l - 1
}

func (t *SmallPayloadNonOptimized) NewJsonMarshal() []byte {
	tmpPtr := lib.GetFromPool()
	tmp := *tmpPtr
	tmp = t.AppendJsonString(tmp)
	res := make([]byte, len(tmp))
	copy(res, tmp)
	*tmpPtr = tmp
	lib.PutToPool(tmpPtr)
	return res
}

//go:nosplit
func (t *SmallPayloadNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"st\":"...)
	res = lib.AppendInt(res, t.St)
	res = append(res, ",\"sid\":"...)
	res = lib.AppendInt(res, t.Sid)
	res = append(res, ",\"tt\":\""...)
	res = lib.AppendByteWithEscape(res, t.Tt)
	res = append(res, "\",\"gr\":"...)
	res = lib.AppendInt(res, t.Gr)
	res = append(res, ",\"uuid\":\""...)
	res = lib.AppendByteWithEscape(res, t.Uuid)
	res = append(res, "\",\"ip\":\""...)
	res = lib.AppendByteWithEscape(res, t.Ip)
	res = append(res, "\",\"ua\":\""...)
	res = lib.AppendByteWithEscape(res, t.Ua)
	res = append(res, "\",\"tz\":"...)
	res = lib.AppendInt(res, t.Tz)
	res = append(res, ",\"v\":"...)
	res = lib.AppendInt(res, t.V)
	res = append(res, '}')
	return res
}

//go:nosplit
func (t *SmallPayloadNonOptimized) JsonLen() int {
	l := 2
	l += 20
	l += 2 + 2 + 1 + 1
	l += 20
	l += 2 + 3 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Tt)
	l += 2 + 2 + 1 + 1
	l += 20
	l += 2 + 2 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Uuid)
	l += 2 + 4 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Ip)
	l += 2 + 2 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Ua)
	l += 2 + 2 + 1 + 1
	l += 20
	l += 2 + 2 + 1 + 1
	l += 20
	l += 2 + 1 + 1 + 1
	return l - 1
}

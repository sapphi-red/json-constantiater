// Code generated by json-constantiater DO NOT EDIT

package benchmark

import "github.com/sapphi-red/json-constantiater/lib"

func (t *DSUser) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSUser) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"username\":\""...)
	res = append(res, t.Username...)
	res = append(res, "\","...)
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *DSUser) JsonLen() int {
	l := 2
	l += 2 + len(t.Username)
	l += 2 + 8 + 1 + 1
	return l - 1
}

func (t *DSTopic) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSTopic) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"id\":"...)
	res = lib.AppendInt(res, t.Id)
	res = append(res, ",\"slug\":\""...)
	res = append(res, t.Slug...)
	res = append(res, "\","...)
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *DSTopic) JsonLen() int {
	l := 2
	l += 20
	l += 2 + 2 + 1 + 1
	l += 2 + len(t.Slug)
	l += 2 + 4 + 1 + 1
	return l - 1
}

func (t *DSTopics) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSTopics) AppendJsonString(res []byte) []byte {
	res = append(res, '[')

	for _, e := range *t {
		if e == nil {
			res = append(res, `null`...)
		} else {
			res = e.AppendJsonString(res)
		}
		res = append(res, ',')
	}
	res[len(res)-1] = ']'
	return res
}

//go:nosplit
func (t *DSTopics) JsonLen() int {
	l := 2
	for _, e := range *t {
		if e == nil {
			l += 4
		} else {
			l += e.JsonLen()
		}
		l += 1
	}
	return l - 1
}

//go:nosplit
func (t *DSTopics) IsEmpty() bool {
	return len(*t) == 0
}

func (t *DSTopicsList) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSTopicsList) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"topics\":"...)
	res = t.Topics.AppendJsonString(res)
	res = append(res, ",\"more_topics_url\":\""...)
	res = append(res, t.MoreTopicsUrl...)
	res = append(res, "\","...)
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *DSTopicsList) JsonLen() int {
	l := 2
	l += t.Topics.JsonLen()
	l += 2 + 6 + 1 + 1
	l += 2 + len(t.MoreTopicsUrl)
	l += 2 + 15 + 1 + 1
	return l - 1
}

func (t *DSUsers) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSUsers) AppendJsonString(res []byte) []byte {
	res = append(res, '[')

	for _, e := range *t {
		if e == nil {
			res = append(res, `null`...)
		} else {
			res = e.AppendJsonString(res)
		}
		res = append(res, ',')
	}
	res[len(res)-1] = ']'
	return res
}

//go:nosplit
func (t *DSUsers) JsonLen() int {
	l := 2
	for _, e := range *t {
		if e == nil {
			l += 4
		} else {
			l += e.JsonLen()
		}
		l += 1
	}
	return l - 1
}

//go:nosplit
func (t *DSUsers) IsEmpty() bool {
	return len(*t) == 0
}

func (t *LargePayload) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *LargePayload) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"users\":"...)
	res = t.Users.AppendJsonString(res)
	res = append(res, ",\"topics\":"...)
	if t.Topics == nil {
		res = append(res, `null`...)
	} else {
		res = t.Topics.AppendJsonString(res)
	}
	res = append(res, ',')
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *LargePayload) JsonLen() int {
	l := 2
	l += t.Users.JsonLen()
	l += 2 + 5 + 1 + 1
	if t.Topics == nil {
		l += 4
	} else {
		l += t.Topics.JsonLen()
	}
	l += 2 + 6 + 1 + 1
	return l - 1
}

func (t *DSUserNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSUserNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"username\":\""...)
	res = lib.AppendByteWithEscape(res, t.Username)
	res = append(res, "\","...)
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *DSUserNonOptimized) JsonLen() int {
	l := 2
	l += 2 + lib.GetEscapedLen(t.Username)
	l += 2 + 8 + 1 + 1
	return l - 1
}

func (t *DSTopicNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSTopicNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"id\":"...)
	res = lib.AppendInt(res, t.Id)
	res = append(res, ",\"slug\":\""...)
	res = lib.AppendByteWithEscape(res, t.Slug)
	res = append(res, "\","...)
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *DSTopicNonOptimized) JsonLen() int {
	l := 2
	l += 20
	l += 2 + 2 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Slug)
	l += 2 + 4 + 1 + 1
	return l - 1
}

func (t *DSTopicsNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSTopicsNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, '[')

	for _, e := range *t {
		if e == nil {
			res = append(res, `null`...)
		} else {
			res = e.AppendJsonString(res)
		}
		res = append(res, ',')
	}
	res[len(res)-1] = ']'
	return res
}

//go:nosplit
func (t *DSTopicsNonOptimized) JsonLen() int {
	l := 2
	for _, e := range *t {
		if e == nil {
			l += 4
		} else {
			l += e.JsonLen()
		}
		l += 1
	}
	return l - 1
}

//go:nosplit
func (t *DSTopicsNonOptimized) IsEmpty() bool {
	return len(*t) == 0
}

func (t *DSTopicsListNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSTopicsListNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"topics\":"...)
	res = t.Topics.AppendJsonString(res)
	res = append(res, ",\"more_topics_url\":\""...)
	res = lib.AppendByteWithEscape(res, t.MoreTopicsUrl)
	res = append(res, "\","...)
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *DSTopicsListNonOptimized) JsonLen() int {
	l := 2
	l += t.Topics.JsonLen()
	l += 2 + 6 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.MoreTopicsUrl)
	l += 2 + 15 + 1 + 1
	return l - 1
}

func (t *DSUsersNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *DSUsersNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, '[')

	for _, e := range *t {
		if e == nil {
			res = append(res, `null`...)
		} else {
			res = e.AppendJsonString(res)
		}
		res = append(res, ',')
	}
	res[len(res)-1] = ']'
	return res
}

//go:nosplit
func (t *DSUsersNonOptimized) JsonLen() int {
	l := 2
	for _, e := range *t {
		if e == nil {
			l += 4
		} else {
			l += e.JsonLen()
		}
		l += 1
	}
	return l - 1
}

//go:nosplit
func (t *DSUsersNonOptimized) IsEmpty() bool {
	return len(*t) == 0
}

func (t *LargePayloadNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *LargePayloadNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"users\":"...)
	res = t.Users.AppendJsonString(res)
	res = append(res, ",\"topics\":"...)
	if t.Topics == nil {
		res = append(res, `null`...)
	} else {
		res = t.Topics.AppendJsonString(res)
	}
	res = append(res, ',')
	res[len(res)-1] = '}'
	return res
}

//go:nosplit
func (t *LargePayloadNonOptimized) JsonLen() int {
	l := 2
	l += t.Users.JsonLen()
	l += 2 + 5 + 1 + 1
	if t.Topics == nil {
		l += 4
	} else {
		l += t.Topics.JsonLen()
	}
	l += 2 + 6 + 1 + 1
	return l - 1
}

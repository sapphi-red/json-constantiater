// Code generated by json-constantiater DO NOT EDIT

package benchmark

import "github.com/sapphi-red/json-constantiater/lib"

//go:nosplit
func (t *CBAvatar) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBAvatar) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"url\":\""...)
	res = append(res, t.Url...)
	res = append(res, '"')

	return append(res, '}')
}

//go:nosplit
func (t *CBAvatar) JsonLen() uint64 {
	var l uint64 = 2
	l += 2 + uint64(len(t.Url))
	l += 2 + 3 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *Avatars) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *Avatars) AppendJsonString(res []byte) []byte {
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
func (t *Avatars) JsonLen() uint64 {
	var l uint64 = 2
	for _, e := range *t {
		l += e.JsonLen() + 1
	}
	return l - 1
}

//go:nosplit
func (t *CBGravatar) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBGravatar) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"avatars\":"...)
	res = t.Avatars.AppendJsonString(res)
	return append(res, '}')
}

//go:nosplit
func (t *CBGravatar) JsonLen() uint64 {
	var l uint64 = 2
	l += t.Avatars.JsonLen()
	l += 2 + 7 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *CBGithub) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBGithub) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"followers\":"...)
	res = lib.AppendInt(res, t.Followers)
	return append(res, '}')
}

//go:nosplit
func (t *CBGithub) JsonLen() uint64 {
	var l uint64 = 2
	l += 20
	l += 2 + 9 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *CBName) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBName) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"fullName\":\""...)
	res = append(res, t.FullName...)
	res = append(res, '"')

	return append(res, '}')
}

//go:nosplit
func (t *CBName) JsonLen() uint64 {
	var l uint64 = 2
	l += 2 + uint64(len(t.FullName))
	l += 2 + 8 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *CBPerson) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBPerson) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"name\":"...)
	if t.Name == nil {
		res = append(res, `null`...)
	} else {
		res = t.Name.AppendJsonString(res)
	}
	res = append(res, ",\"github\":"...)
	if t.Github == nil {
		res = append(res, `null`...)
	} else {
		res = t.Github.AppendJsonString(res)
	}
	res = append(res, ",\"gravatar\":"...)
	if t.Gravatar == nil {
		res = append(res, `null`...)
	} else {
		res = t.Gravatar.AppendJsonString(res)
	}
	return append(res, '}')
}

//go:nosplit
func (t *CBPerson) JsonLen() uint64 {
	var l uint64 = 2
	if t.Name == nil {
		l += 4
	} else {
		l += t.Name.JsonLen()
	}
	l += 2 + 4 + 1 + 1
	if t.Github == nil {
		l += 4
	} else {
		l += t.Github.JsonLen()
	}
	l += 2 + 6 + 1 + 1
	if t.Gravatar == nil {
		l += 4
	} else {
		l += t.Gravatar.JsonLen()
	}
	l += 2 + 8 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *MediumPayload) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *MediumPayload) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"person\":"...)
	if t.Person == nil {
		res = append(res, `null`...)
	} else {
		res = t.Person.AppendJsonString(res)
	}
	res = append(res, ",\"company\":\""...)
	res = lib.AppendByteWithEscape(res, t.Company)
	res = append(res, '"')

	return append(res, '}')
}

//go:nosplit
func (t *MediumPayload) JsonLen() uint64 {
	var l uint64 = 2
	if t.Person == nil {
		l += 4
	} else {
		l += t.Person.JsonLen()
	}
	l += 2 + 6 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Company)
	l += 2 + 7 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *CBAvatarNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBAvatarNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"url\":\""...)
	res = lib.AppendByteWithEscape(res, t.Url)
	res = append(res, '"')

	return append(res, '}')
}

//go:nosplit
func (t *CBAvatarNonOptimized) JsonLen() uint64 {
	var l uint64 = 2
	l += 2 + lib.GetEscapedLen(t.Url)
	l += 2 + 3 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *AvatarsNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *AvatarsNonOptimized) AppendJsonString(res []byte) []byte {
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
func (t *AvatarsNonOptimized) JsonLen() uint64 {
	var l uint64 = 2
	for _, e := range *t {
		l += e.JsonLen() + 1
	}
	return l - 1
}

//go:nosplit
func (t *CBGravatarNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBGravatarNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"avatars\":"...)
	res = t.Avatars.AppendJsonString(res)
	return append(res, '}')
}

//go:nosplit
func (t *CBGravatarNonOptimized) JsonLen() uint64 {
	var l uint64 = 2
	l += t.Avatars.JsonLen()
	l += 2 + 7 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *CBGithubNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBGithubNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"followers\":"...)
	res = lib.AppendInt(res, t.Followers)
	return append(res, '}')
}

//go:nosplit
func (t *CBGithubNonOptimized) JsonLen() uint64 {
	var l uint64 = 2
	l += 20
	l += 2 + 9 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *CBNameNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBNameNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"fullName\":\""...)
	res = lib.AppendByteWithEscape(res, t.FullName)
	res = append(res, '"')

	return append(res, '}')
}

//go:nosplit
func (t *CBNameNonOptimized) JsonLen() uint64 {
	var l uint64 = 2
	l += 2 + lib.GetEscapedLen(t.FullName)
	l += 2 + 8 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *CBPersonNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *CBPersonNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"name\":"...)
	if t.Name == nil {
		res = append(res, `null`...)
	} else {
		res = t.Name.AppendJsonString(res)
	}
	res = append(res, ",\"github\":"...)
	if t.Github == nil {
		res = append(res, `null`...)
	} else {
		res = t.Github.AppendJsonString(res)
	}
	res = append(res, ",\"gravatar\":"...)
	if t.Gravatar == nil {
		res = append(res, `null`...)
	} else {
		res = t.Gravatar.AppendJsonString(res)
	}
	return append(res, '}')
}

//go:nosplit
func (t *CBPersonNonOptimized) JsonLen() uint64 {
	var l uint64 = 2
	if t.Name == nil {
		l += 4
	} else {
		l += t.Name.JsonLen()
	}
	l += 2 + 4 + 1 + 1
	if t.Github == nil {
		l += 4
	} else {
		l += t.Github.JsonLen()
	}
	l += 2 + 6 + 1 + 1
	if t.Gravatar == nil {
		l += 4
	} else {
		l += t.Gravatar.JsonLen()
	}
	l += 2 + 8 + 1 + 1
	return l - 1
}

//go:nosplit
func (t *MediumPayloadNonOptimized) NewJsonMarshal() []byte {
	res := make([]byte, 0, t.JsonLen())
	return t.AppendJsonString(res)
}

//go:nosplit
func (t *MediumPayloadNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"person\":"...)
	if t.Person == nil {
		res = append(res, `null`...)
	} else {
		res = t.Person.AppendJsonString(res)
	}
	res = append(res, ",\"company\":\""...)
	res = lib.AppendByteWithEscape(res, t.Company)
	res = append(res, '"')

	return append(res, '}')
}

//go:nosplit
func (t *MediumPayloadNonOptimized) JsonLen() uint64 {
	var l uint64 = 2
	if t.Person == nil {
		l += 4
	} else {
		l += t.Person.JsonLen()
	}
	l += 2 + 6 + 1 + 1
	l += 2 + lib.GetEscapedLen(t.Company)
	l += 2 + 7 + 1 + 1
	return l - 1
}

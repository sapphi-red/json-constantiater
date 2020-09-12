// Code generated by json-constantiater DO NOT EDIT

package benchmark

import "github.com/sapphi-red/json-constantiater/lib"

func (t *CBAvatar) NewJsonMarshal() []byte {
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
func (t *CBAvatar) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"url\":\""...)
	res = append(res, t.Url...)
	res = append(res, "\"}"...)
	return res
}

func (t *Avatars) NewJsonMarshal() []byte {
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
func (t *Avatars) AppendJsonString(res []byte) []byte {
	if len(*t) <= 0 {
		return append(res, `[]`...)
	}
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
func (t *Avatars) IsEmpty() bool {
	return len(*t) == 0
}

func (t *CBGravatar) NewJsonMarshal() []byte {
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
func (t *CBGravatar) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"avatars\":"...)
	res = t.Avatars.AppendJsonString(res)
	res = append(res, '}')
	return res
}

func (t *CBGithub) NewJsonMarshal() []byte {
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
func (t *CBGithub) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"followers\":"...)
	if t.Followers < lib.NSmalls {
		res = lib.AppendSmallInt(res, t.Followers)
	} else {
		res = lib.AppendInt(res, t.Followers)
	}
	res = append(res, '}')
	return res
}

func (t *CBName) NewJsonMarshal() []byte {
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
func (t *CBName) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"fullName\":\""...)
	res = append(res, t.FullName...)
	res = append(res, "\"}"...)
	return res
}

func (t *CBPerson) NewJsonMarshal() []byte {
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
	res = append(res, '}')
	return res
}

func (t *MediumPayload) NewJsonMarshal() []byte {
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
func (t *MediumPayload) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"person\":"...)
	if t.Person == nil {
		res = append(res, `null`...)
	} else {
		res = t.Person.AppendJsonString(res)
	}
	res = append(res, ",\"company\":\""...)
	res = lib.AppendByteWithEscape(res, t.Company)
	res = append(res, "\"}"...)
	return res
}

func (t *CBAvatarNonOptimized) NewJsonMarshal() []byte {
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
func (t *CBAvatarNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"url\":\""...)
	res = lib.AppendByteWithEscape(res, t.Url)
	res = append(res, "\"}"...)
	return res
}

func (t *AvatarsNonOptimized) NewJsonMarshal() []byte {
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
func (t *AvatarsNonOptimized) AppendJsonString(res []byte) []byte {
	if len(*t) <= 0 {
		return append(res, `[]`...)
	}
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
func (t *AvatarsNonOptimized) IsEmpty() bool {
	return len(*t) == 0
}

func (t *CBGravatarNonOptimized) NewJsonMarshal() []byte {
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
func (t *CBGravatarNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"avatars\":"...)
	res = t.Avatars.AppendJsonString(res)
	res = append(res, '}')
	return res
}

func (t *CBGithubNonOptimized) NewJsonMarshal() []byte {
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
func (t *CBGithubNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"followers\":"...)
	if 0 <= t.Followers {
		if t.Followers < lib.NSmalls {
			res = lib.AppendSmallInt(res, t.Followers)
		} else {
			res = lib.AppendInt(res, t.Followers)
		}
	} else {
		if -lib.NSmalls < t.Followers {
			res = lib.AppendSmallMinusInt(res, t.Followers)
		} else {
			res = lib.AppendInt(res, t.Followers)
		}
	}
	res = append(res, '}')
	return res
}

func (t *CBNameNonOptimized) NewJsonMarshal() []byte {
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
func (t *CBNameNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"fullName\":\""...)
	res = lib.AppendByteWithEscape(res, t.FullName)
	res = append(res, "\"}"...)
	return res
}

func (t *CBPersonNonOptimized) NewJsonMarshal() []byte {
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
	res = append(res, '}')
	return res
}

func (t *MediumPayloadNonOptimized) NewJsonMarshal() []byte {
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
func (t *MediumPayloadNonOptimized) AppendJsonString(res []byte) []byte {
	res = append(res, "{\"person\":"...)
	if t.Person == nil {
		res = append(res, `null`...)
	} else {
		res = t.Person.AppendJsonString(res)
	}
	res = append(res, ",\"company\":\""...)
	res = lib.AppendByteWithEscape(res, t.Company)
	res = append(res, "\"}"...)
	return res
}

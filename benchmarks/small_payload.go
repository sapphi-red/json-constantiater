package benchmark

import "github.com/francoispqt/gojay"

//easyjson:json
type SmallPayload struct {
	St   int `json:",unsigned"`
	Sid  int `json:",unsigned"`
	Tt   string
	Gr   int
	Uuid string `json:",noescape"`
	Ip   string `json:",noescape"`
	Ua   string
	Tz   int `json:",small"`
	V    int `json:",unsigned,small"`
}

func (t *SmallPayload) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddIntKey("st", t.St)
	enc.AddIntKey("sid", t.Sid)
	enc.AddStringKey("tt", t.Tt)
	enc.AddIntKey("gr", t.Gr)
	enc.AddStringKey("uuid", t.Uuid)
	enc.AddStringKey("ip", t.Ip)
	enc.AddStringKey("ua", t.Ua)
	enc.AddIntKey("tz", t.Tz)
	enc.AddIntKey("v", t.V)
}

func (t *SmallPayload) IsNil() bool {
	return t == nil
}

func NewSmallPayload() *SmallPayload {
	return &SmallPayload{
		St:   1,
		Sid:  2,
		Tt:   "TestString",
		Gr:   4,
		Uuid: "8f9a65eb-4807-4d57-b6e0-bda5d62f1429",
		Ip:   "127.0.0.1",
		Ua:   "Mozilla",
		Tz:   8,
		V:    6,
	}
}

// -----

type SmallPayloadNonOptimized struct {
	St   int
	Sid  int
	Tt   string
	Gr   int
	Uuid string
	Ip   string
	Ua   string
	Tz   int
	V    int
}

func NewSmallPayloadNonOptimized() *SmallPayloadNonOptimized {
	return &SmallPayloadNonOptimized{
		St:   1,
		Sid:  2,
		Tt:   "TestString",
		Gr:   4,
		Uuid: "8f9a65eb-4807-4d57-b6e0-bda5d62f1429",
		Ip:   "127.0.0.1",
		Ua:   "Mozilla",
		Tz:   8,
		V:    6,
	}
}

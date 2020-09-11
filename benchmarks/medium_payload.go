package benchmark

import "github.com/francoispqt/gojay"

type CBAvatar struct {
	Url string `json:",noescape"`
}

func (m *CBAvatar) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddStringKey("url", m.Url)
}

func (m *CBAvatar) IsNil() bool {
	return m == nil
}

type Avatars []*CBAvatar

func (m *Avatars) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *m {
		enc.AddObject(e)
	}
}
func (m *Avatars) IsNil() bool {
	return m == nil
}

type CBGravatar struct {
	Avatars Avatars
}

func (m *CBGravatar) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey("avatars", &m.Avatars)
}

func (m *CBGravatar) IsNil() bool {
	return m == nil
}

type CBGithub struct {
	Followers int
}

func (m *CBGithub) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddIntKey("followers", m.Followers)
}

func (m *CBGithub) IsNil() bool {
	return m == nil
}

type CBName struct {
	FullName string `json:"fullName,noescape"`
}

func (m *CBName) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddStringKey("fullName", m.FullName)
}

func (m *CBName) IsNil() bool {
	return m == nil
}

type CBPerson struct {
	Name     *CBName   `json:"name"`
	Github   *CBGithub `json:"github"`
	Gravatar *CBGravatar
}

func (m *CBPerson) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddObjectKey("name", m.Name)
	enc.AddObjectKey("github", m.Github)
	enc.AddObjectKey("gravatar", m.Gravatar)
}

func (m *CBPerson) IsNil() bool {
	return m == nil
}

//easyjson:json
type MediumPayload struct {
	Person  *CBPerson `json:"person"`
	Company string    `json:"company"`
}

func (m *MediumPayload) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddObjectKey("person", m.Person)
	enc.AddStringKey("company", m.Company)
}

func (m *MediumPayload) IsNil() bool {
	return m == nil
}

func NewMediumPayload() *MediumPayload {
	return &MediumPayload{
		Company: "test",
		Person: &CBPerson{
			Name: &CBName{
				FullName: "test",
			},
			Github: &CBGithub{
				Followers: 100,
			},
			Gravatar: &CBGravatar{
				Avatars: Avatars{
					&CBAvatar{
						Url: "http://test.com",
					},
					&CBAvatar{
						Url: "http://test.com",
					},
					&CBAvatar{
						Url: "http://test.com",
					},
					&CBAvatar{
						Url: "http://test.com",
					},
					&CBAvatar{
						Url: "http://test.com",
					},
					&CBAvatar{
						Url: "http://test.com",
					},
					&CBAvatar{
						Url: "http://test.com",
					},
					&CBAvatar{
						Url: "http://test.com",
					},
				},
			},
		},
	}
}

// -----

type CBAvatarNonOptimized struct {
	Url string
}

type AvatarsNonOptimized []*CBAvatarNonOptimized

type CBGravatarNonOptimized struct {
	Avatars AvatarsNonOptimized
}

type CBGithubNonOptimized struct {
	Followers int
}

type CBNameNonOptimized struct {
	FullName string `json:"fullName"`
}

type CBPersonNonOptimized struct {
	Name     *CBNameNonOptimized   `json:"name"`
	Github   *CBGithubNonOptimized `json:"github"`
	Gravatar *CBGravatarNonOptimized
}

type MediumPayloadNonOptimized struct {
	Person  *CBPersonNonOptimized `json:"person"`
	Company string                `json:"company"`
}

func NewMediumPayloadNonOptimized() *MediumPayloadNonOptimized {
	return &MediumPayloadNonOptimized{
		Company: "test",
		Person: &CBPersonNonOptimized{
			Name: &CBNameNonOptimized{
				FullName: "test",
			},
			Github: &CBGithubNonOptimized{
				Followers: 100,
			},
			Gravatar: &CBGravatarNonOptimized{
				Avatars: AvatarsNonOptimized{
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
					&CBAvatarNonOptimized{
						Url: "http://test.com",
					},
				},
			},
		},
	}
}

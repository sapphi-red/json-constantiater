package benchmark

import (
	"strconv"

	"github.com/francoispqt/gojay"
)

type DSUser struct {
	Username string `json:",noescape"`
}

func (m *DSUser) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddStringKey("username", m.Username)
}
func (m *DSUser) IsNil() bool {
	return m == nil
}

type DSTopic struct {
	Id   int    `json:",unsigned"`
	Slug string `json:",noescape"`
}

func (m *DSTopic) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddIntKey("id", m.Id)
	enc.AddStringKey("slug", m.Slug)
}
func (m *DSTopic) IsNil() bool {
	return m == nil
}

//value:",nonnil"
type DSTopics []*DSTopic

func (m *DSTopics) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *m {
		enc.AddObject(e)
	}
}
func (m *DSTopics) IsNil() bool {
	return m == nil
}

type DSTopicsList struct {
	Topics        DSTopics
	MoreTopicsUrl string `json:",noescape"`
}

func (m *DSTopicsList) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey("users", &m.Topics)
	enc.AddStringKey("more_topics_url", m.MoreTopicsUrl)
}
func (m *DSTopicsList) IsNil() bool {
	return m == nil
}

//value:",nonnil"
type DSUsers []*DSUser

func (m *DSUsers) MarshalJSONArray(enc *gojay.Encoder) {
	for _, e := range *m {
		enc.AddObject(e)
	}
}
func (m *DSUsers) IsNil() bool {
	return m == nil
}

//easyjson:json
type LargePayload struct {
	Users  DSUsers
	Topics *DSTopicsList `json:",nonnil"`
}

func (m *LargePayload) MarshalJSONObject(enc *gojay.Encoder) {
	enc.AddArrayKey("users", &m.Users)
	enc.AddObjectKey("topics", m.Topics)
}

func (m *LargePayload) IsNil() bool {
	return m == nil
}

func NewLargePayload() *LargePayload {
	dsUsers := DSUsers{}
	dsTopics := DSTopics{}
	for i := 0; i < 100; i++ {
		str := "test" + strconv.Itoa(i)
		dsUsers = append(
			dsUsers,
			&DSUser{
				Username: str,
			},
		)
		dsTopics = append(
			dsTopics,
			&DSTopic{
				Id:   i,
				Slug: str,
			},
		)
	}
	return &LargePayload{
		Users: dsUsers,
		Topics: &DSTopicsList{
			Topics:        dsTopics,
			MoreTopicsUrl: "http://test.com",
		},
	}
}

// -----

type DSUserNonOptimized struct {
	Username string
}

type DSTopicNonOptimized struct {
	Id   int
	Slug string
}

type DSTopicsNonOptimized []*DSTopicNonOptimized

type DSTopicsListNonOptimized struct {
	Topics        DSTopicsNonOptimized
	MoreTopicsUrl string
}

type DSUsersNonOptimized []*DSUserNonOptimized

type LargePayloadNonOptimized struct {
	Users  DSUsersNonOptimized
	Topics *DSTopicsListNonOptimized
}

func NewLargePayloadNonOptimized() *LargePayloadNonOptimized {
	dsUsers := DSUsersNonOptimized{}
	dsTopics := DSTopicsNonOptimized{}
	for i := 0; i < 100; i++ {
		str := "test" + strconv.Itoa(i)
		dsUsers = append(
			dsUsers,
			&DSUserNonOptimized{
				Username: str,
			},
		)
		dsTopics = append(
			dsTopics,
			&DSTopicNonOptimized{
				Id:   i,
				Slug: str,
			},
		)
	}
	return &LargePayloadNonOptimized{
		Users: dsUsers,
		Topics: &DSTopicsListNonOptimized{
			Topics:        dsTopics,
			MoreTopicsUrl: "http://test.com",
		},
	}
}

package event

import (
	"encoding/json"
	"fmt"
	"mockers/pkg/hash"
	"mockers/pkg/utils"
	"time"
)

// NewModel return a Model, and set CreateTime is now.
func NewModel(action, origin, info string, labels map[string]string) *Model {
	var m = &Model{
		Action:     action,
		Origin:     origin,
		Info:       info,
		Labels:     hash.SyncTags{},
		CreateTime: time.Now().UnixNano(),
	}
	m.Labels.FromMap(labels)
	return m
}

// Model is an event object, it contains the Action, Origin, Info, CreateTime.
// The has only cal the Action, Origin, Info and CreateTime.
type Model struct {

	// Action of event Model is define the action. In http request, it can use as method.
	Action string `json:"action"`

	// Origin is the event creator.
	Origin string `json:"origin"`

	// Info is what happen in event.
	Info string `json:"info"`

	// Labels is the event label, the subscribes can listen the label.
	Labels hash.SyncTags `json:"labels"`

	// CreateTime is the event Model create time. Unit is a nano-second.
	CreateTime int64 `json:"create_time"`
}

func (m *Model) Log() string {
	return fmt.Sprintf("Action: [%s], Origin: [%s], Info: [%s], Hash: [%s], CreateTime: [%d], Labels: [%v]",
		m.Action, m.Origin, m.Info, m.Hash(), m.CreateTime, m.Labels.Map())
}

func (m *Model) Hash() string {
	aoi := m.Action + m.Origin + m.Info
	labelHash := m.Labels.Hash()
	cth := utils.Int64ToBytes(m.CreateTime)
	rs := append(cth, []byte(aoi+labelHash)...)
	return hash.SHA256(rs)
}

func (m *Model) IsHash(v string) bool {
	return m.Hash() == v
}

func (m *Model) Serialization() []byte {
	bytes, _ := json.MarshalIndent(m, "", "")
	return bytes
}

func (m *Model) Deserialization(in []byte) error {
	return json.Unmarshal(in, m)
}

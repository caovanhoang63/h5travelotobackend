package chatmodel

import (
	"h5travelotobackend/common"
	"time"
)

type Message struct {
	Id         int           `json:"id" bson:"id"`
	Message    string        `json:"message,omitempty" bson:"message,omitempty"`
	Image      *common.Image `json:"image,omitempty" bson:"image,omitempty"`
	From       int           `json:"-" bson:"from"`
	FromFakeId *common.UID   `json:"from" bson:"-"`
	Status     int           `json:"status" bson:"status"`
	CreatedAt  *time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt  *time.Time    `json:"updated_at" bson:"updated_at"`
}

type Messages []*Message

func (m *Message) OnCreate() {
	newTime := time.Now()
	m.CreatedAt = &newTime
	m.UpdatedAt = &newTime
	m.Status = 1
}

func (m *Message) OnUpdate() {
	newTime := time.Now()
	m.UpdatedAt = &newTime
}

func (m *Message) Mask(isAdmin bool) {
	m.FromFakeId = common.NewUIDP(uint32(m.From), common.DbTypeUser, 0)
}

func (m *Messages) Mask() {
	if m == nil {
		return
	}
	for i := range *m {
		(*m)[i].Mask(false)
	}
}

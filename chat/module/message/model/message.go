package chatmessagemodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"h5travelotobackend/common"
)

const EntityName = "Message"

type Message struct {
	common.MongoModel `bson:",inline"`
	RoomId            *primitive.ObjectID `json:"room_id" bson:"room_id"`
	Message           string              `json:"message,omitempty" bson:"message,omitempty"`
	Image             *common.Image       `json:"image,omitempty" bson:"image,omitempty"`
	UserId            int                 `json:"-" bson:"from"`
	UserFakeId        *common.UID         `json:"from" bson:"-"`
	IsRead            bool                `json:"is_read" bson:"is_read"`
}

func (m *Message) GetChatMessageId() string {
	return m.ID.Hex()
}

func (m *Message) GetChatRoomId() string {
	return m.RoomId.Hex()
}

func (m *Message) GetSenderId() int {
	return m.UserId
}

func (m Message) CollectionName() string {
	return "chat_messages"
}

type Messages []*Message

func (m *Message) Mask(isAdmin bool) {
	m.UserFakeId = common.NewUIDP(uint32(m.
		UserId), common.DbTypeUser, 0)
}

func (m *Messages) Mask() {
	if m == nil {
		return
	}
	for i := range *m {
		(*m)[i].Mask(false)
	}
}

package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoModel struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status    int                 `json:"status" bson:"status,omitempty" default:"1"`
	CreatedAt *time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time          `json:"updated_at" bson:"updated_at"`
}

func (m *MongoModel) FixTime() {
	if m.CreatedAt != nil {
		*m.CreatedAt = m.CreatedAt.Add(time.Hour * 7)
	}
	if m.UpdatedAt != nil {
		*m.UpdatedAt = m.UpdatedAt.Add(time.Hour * 7)
	}
}

func (m *MongoModel) OnCreate() {
	m.Status = 1
	newTime := time.Now()
	m.CreatedAt = &newTime
	m.UpdatedAt = &newTime
}

func (m *MongoModel) OnUpdate() {
	newTime := time.Now()
	m.CreatedAt = &newTime
	m.UpdatedAt = &newTime
}

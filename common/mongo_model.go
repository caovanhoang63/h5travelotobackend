package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoModel struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Status    int                `json:"status" bson:"status,,omitempty" default:"1"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (m *MongoModel) OnCreate() {
	m.Status = 1
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

func (m *MongoModel) OnUpdate() {
	m.UpdatedAt = time.Now()
}

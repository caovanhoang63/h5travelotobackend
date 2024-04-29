package common

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"h5travelotobackend/component/pubsub"
)

type SimpleAppContext interface {
	GetGormDbConnection() *gorm.DB
	GetSecretKey() string
	GetMongoConnection() *mongo.Database
	GetPubSub() pubsub.Pubsub
}

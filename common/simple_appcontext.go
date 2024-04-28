package common

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type SimpleAppContext interface {
	GetGormDbConnection() *gorm.DB
	GetSecretKey() string
	GetMongoConnection() *mongo.Database
}

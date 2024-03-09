package appContext

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type AppContext interface {
	GetGormDbConnection() *gorm.DB
	GetMongoConnection() *mongo.Database
	GetSecretKey() string
}

type appContext struct {
	db        *gorm.DB
	mongodb   *mongo.Database
	secretKey string
}

func NewAppContext(db *gorm.DB, mongodb *mongo.Database, secretKey string) AppContext {
	return &appContext{
		db:        db,
		mongodb:   mongodb,
		secretKey: secretKey,
	}
}

func (appCtx *appContext) GetGormDbConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) GetMongoConnection() *mongo.Database {
	return appCtx.mongodb
}

func (appCtx *appContext) GetSecretKey() string { return appCtx.secretKey }

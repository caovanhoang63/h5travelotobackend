package appContext

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type AppContext interface {
	GetGormDbConnection() *gorm.DB
	GetMongoConnection() *mongo.Database
}

type appContext struct {
	db      *gorm.DB
	mongodb *mongo.Database
}

func NewAppContext(db *gorm.DB, mongodb *mongo.Database) AppContext {
	return &appContext{
		db:      db,
		mongodb: mongodb,
	}
}

func (appCtx *appContext) GetGormDbConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) GetMongoConnection() *mongo.Database {
	return appCtx.mongodb
}

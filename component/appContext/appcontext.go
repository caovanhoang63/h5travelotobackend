package appContext

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"h5travelotobackend/component/pubsub"
	"h5travelotobackend/component/uploadprovider"
)

type AppContext interface {
	GetGormDbConnection() *gorm.DB
	GetMongoConnection() *mongo.Database
	GetSecretKey() string
	UploadProvider() uploadprovider.UploadProvider
	GetPubSub() pubsub.Pubsub
}

type appContext struct {
	db             *gorm.DB
	mongodb        *mongo.Database
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	pubSub         pubsub.Pubsub
}

func NewAppContext(db *gorm.DB, mongodb *mongo.Database, secretKey string, provider uploadprovider.UploadProvider, pubsub pubsub.Pubsub) AppContext {
	return &appContext{
		db:             db,
		mongodb:        mongodb,
		secretKey:      secretKey,
		uploadProvider: provider,
		pubSub:         pubsub,
	}
}

func (appCtx *appContext) GetGormDbConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) GetMongoConnection() *mongo.Database {
	return appCtx.mongodb
}

func (appCtx *appContext) GetSecretKey() string { return appCtx.secretKey }

func (appCtx *appContext) UploadProvider() uploadprovider.UploadProvider {
	return appCtx.uploadProvider
}

func (appCtx *appContext) GetPubSub() pubsub.Pubsub { return appCtx.pubSub }

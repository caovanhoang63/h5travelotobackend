package appContext

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"h5travelotobackend/component/uploadprovider"
)

type AppContext interface {
	GetGormDbConnection() *gorm.DB
	GetMongoConnection() *mongo.Database
	GetSecretKey() string
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db             *gorm.DB
	mongodb        *mongo.Database
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, mongodb *mongo.Database, secretKey string, provider uploadprovider.UploadProvider) AppContext {
	return &appContext{
		db:             db,
		mongodb:        mongodb,
		secretKey:      secretKey,
		uploadProvider: provider,
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

package appContext

import (
	"github.com/elastic/go-elasticsearch/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"h5travelotobackend/component/pubsub"
	"h5travelotobackend/component/uploadprovider"
	"h5travelotobackend/skio"
)

type AppContext interface {
	GetGormDbConnection() *gorm.DB
	GetMongoConnection() *mongo.Database
	GetSecretKey() string
	UploadProvider() uploadprovider.UploadProvider
	GetPubSub() pubsub.Pubsub
	GetRealTimeEngine() skio.RealtimeEngine
	GetElasticSearchClient() *elasticsearch.TypedClient
}

type appContext struct {
	db             *gorm.DB
	mongodb        *mongo.Database
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	pubSub         pubsub.Pubsub
	rtEngine       skio.RealtimeEngine
	esClient       *elasticsearch.TypedClient
}

func NewAppContext(db *gorm.DB,
	mongodb *mongo.Database,
	secretKey string,
	provider uploadprovider.UploadProvider,
	pubsub pubsub.Pubsub,
	es *elasticsearch.TypedClient) *appContext {
	return &appContext{
		db:             db,
		mongodb:        mongodb,
		secretKey:      secretKey,
		uploadProvider: provider,
		pubSub:         pubsub,
		esClient:       es,
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

func (appContext *appContext) GetRealTimeEngine() skio.RealtimeEngine { return appContext.rtEngine }

func (appContext *appContext) SetRealTimeEngine(rtEngine skio.RealtimeEngine) {
	appContext.rtEngine = rtEngine
}

func (appCtx *appContext) GetElasticSearchClient() *elasticsearch.TypedClient {
	return appCtx.esClient
}

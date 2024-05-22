package appContext

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"h5travelotobackend/component/cacher"
	"h5travelotobackend/component/logger"
	"h5travelotobackend/component/payment/vnpay"
	"h5travelotobackend/component/pubsub"
	"h5travelotobackend/component/uploadprovider"
	"h5travelotobackend/component/uuid"
	"h5travelotobackend/email"
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
	GetRedisClient() *redis.Client
	GetLogger() logger.Logger
	GetCacher() cacher.Cacher
	GetVnPay() *vnpay.VnPay
	GetUUID() uuid.Uuid
	GetSendMailEngine() email.Engine
}

type appContext struct {
	db             *gorm.DB
	mongodb        *mongo.Database
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	pubSub         pubsub.Pubsub
	rtEngine       skio.RealtimeEngine
	esClient       *elasticsearch.TypedClient
	redisClient    *redis.Client
	logger         logger.Logger
	cacher         cacher.Cacher
	vnPay          *vnpay.VnPay
	uuid           uuid.Uuid
	emailEngine    email.Engine
}

func NewAppContext(db *gorm.DB,
	mongodb *mongo.Database,
	secretKey string,
	provider uploadprovider.UploadProvider,
	pubsub pubsub.Pubsub,
	es *elasticsearch.TypedClient,
	redisClient *redis.Client,
	logger logger.Logger,
	cacher cacher.Cacher,
	vnPay *vnpay.VnPay,
	uuid uuid.Uuid,
	emailEngine email.Engine) *appContext {
	return &appContext{
		db:             db,
		mongodb:        mongodb,
		secretKey:      secretKey,
		uploadProvider: provider,
		pubSub:         pubsub,
		esClient:       es,
		redisClient:    redisClient,
		logger:         logger,
		cacher:         cacher,
		vnPay:          vnPay,
		uuid:           uuid,
		emailEngine:    emailEngine,
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

func (appCtx *appContext) GetRedisClient() *redis.Client {
	return appCtx.redisClient
}

func (appCtx *appContext) GetLogger() logger.Logger {
	return appCtx.logger
}

func (appCtx *appContext) GetCacher() cacher.Cacher {
	return appCtx.cacher
}
func (appCtx *appContext) GetVnPay() *vnpay.VnPay {
	return appCtx.vnPay
}
func (appCtx *appContext) GetUUID() uuid.Uuid {
	return appCtx.uuid
}

func (appCtx *appContext) GetSendMailEngine() email.Engine {
	return appCtx.emailEngine
}

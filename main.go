package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"h5travelotobackend/component/appContext"
	rabbitpubsub "h5travelotobackend/component/pubsub/rabbitmq"
	"h5travelotobackend/component/uploadprovider"
	"h5travelotobackend/middleware"
	"h5travelotobackend/skio"
	"h5travelotobackend/subcriber"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	isDev := true

	if isDev {
		err := godotenv.Load(".dev.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	RegisterImageFormat()
	// Get environment variables
	systemSecretKey := os.Getenv("SYSTEM_SECRET_KEY")
	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3ApiKey := os.Getenv("S3_API_KEY")
	s3Secret := os.Getenv("S3_SECRET")
	s3Domain := os.Getenv("S3_DOMAIN")
	mySqlConnString := os.Getenv("MYSQL_CONN_STRING")
	mongoDbConnString := os.Getenv("MONGODB_CONN_STRING")
	rabbitMqConnString := os.Getenv("RABBITMQ_CONN_STRING")
	esURL := os.Getenv("ES_URL")
	esUserName := os.Getenv("ES_USERNAME")
	esPassword := os.Getenv("ES_PASSWORD")
	redisConnString := os.Getenv("REDIS_CONN_STRING")

	// Set up Elasticsearch Connection
	esCfg := elasticsearch.Config{
		Addresses: []string{
			esURL,
		},
		Username: esUserName,
		Password: esPassword,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	es, err := elasticsearch.NewTypedClient(esCfg)
	if err != nil {
		log.Println("Error creating the client: ", err)
	} else {
		log.Println("Elasticsearch client created")
	}

	ping, err := es.Ping().Do(context.Background())
	if err != nil {
		log.Println("Error creating the client: ", err)
	} else {
		log.Println("Elasticsearch ping: ", ping)
	}
	//// End Set up Elasticsearch Connection
	//
	//// Set up Redis Connection
	redisConnOpt, err := redis.ParseURL(redisConnString)
	if err != nil {
		log.Println("Error parsing Redis URL: ", err)
	}
	redisClient := redis.NewClient(redisConnOpt)
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Println("Error connecting to Redis: ", err)
	} else {
		log.Println("Connected to Redis")
	}
	//// End Set up Redis Connection

	// Set up MongoDb Connection
	/***************************************************************/
	/***************************************************************/
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoDbConnString).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Database("h5traveloto").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	mongodb := client.Database("h5traveloto")
	/***************************************************************/
	/***************************************************************/

	// Set up Mysql Gorm Connection
	/***************************************************************/
	/***************************************************************/
	db, err := gorm.Open(mysql.Open(mySqlConnString), &gorm.Config{})
	if err != nil {
		log.Fatal(db, err)
	}
	db = db.Debug()
	/***************************************************************/
	/***************************************************************/

	// Set up S3 Provider
	/***************************************************************/
	/***************************************************************/
	s3Provider := uploadprovider.NewS3Provider(
		s3BucketName,
		s3Region,
		s3ApiKey,
		s3Secret,
		s3Domain,
	)
	/***************************************************************/
	/***************************************************************/

	// Set up RabbitMq Connection
	/***************************************************************/
	/***************************************************************/
	rabbitConn, err := amqp.Dial(rabbitMqConnString)
	if err != nil {
		log.Fatal("Fail to connect rabbitMQ! ", err)
	}
	defer rabbitConn.Close()
	ch, err := rabbitConn.Channel()
	if err != nil {
		log.Fatal("Fail to open channel! ", err)
	}

	pb := rabbitpubsub.NewRabbitPubSub(ch)
	/***************************************************************/
	/***************************************************************/

	// Set up App Context
	appCtx := appContext.NewAppContext(db, mongodb, systemSecretKey, s3Provider, pb, es, redisClient)

	r := gin.New()
	r.Use(middleware.Recover(appCtx))
	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/v1")

	v1.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.OPTIONS("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	SetUpRoute(appCtx, v1)

	err = subcriber.NewEngine(appCtx).Start()
	if err != nil {
		log.Fatal(err)
	}

	rtEngine := skio.NewEngine()
	appCtx.SetRealTimeEngine(rtEngine)
	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Println(err)
	}

	r.StaticFile("/socket/hotel", "./hotel.html")
	r.StaticFile("/socket/customer", "./customer.html")

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

package rdpubsub

import (
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	"log"
)

type redisPubsub struct {
	rdb *redis.Client
}

func NewRedisPubSub(rdb *redis.Client) *redisPubsub {
	return &redisPubsub{
		rdb: rdb,
	}
}

func (pb *redisPubsub) Publish(ctx context.Context, exchange string, data *pubsub.Message) error {
	go func() {
		defer common.AppRecover()
		data.SetChannel(exchange)
		json, err := data.Marshal()
		if err != nil {
			return
		}

		pb.rdb.Publish(ctx, exchange, json)

		log.Println("New event published:", exchange, "with data", string(data.Data))
	}()
	return nil
}

func (pb *redisPubsub) Subscribe(ctx context.Context, exchange string) (<-chan *pubsub.Message, func()) {
	rdpb := pb.rdb.Subscribe(ctx, exchange)

	err := rdpb.Ping(ctx)
	if err != nil {
		failOnError(err, "Failed to register a consumer")
	}

	ch := make(chan *pubsub.Message)

	go func() {
		defer rdpb.Close()
		for msg := range rdpb.Channel() {
			var mess pubsub.Message
			mess.Unmarshal([]byte(msg.Payload))
			ch <- &mess
		}
	}()
	return ch, func() {
		return
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

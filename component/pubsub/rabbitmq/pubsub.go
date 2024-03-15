package rabbitpubsub

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	"log"
	"sync"
)

type rabbitPubSub struct {
	channel    *amqp.Channel
	queue      *amqp.Queue
	mapChannel map[string][]chan *pubsub.Message
	locker     *sync.RWMutex
}

func NewRabbitPubSub(channel *amqp.Channel) *rabbitPubSub {
	return &rabbitPubSub{
		channel:    channel,
		mapChannel: make(map[string][]chan *pubsub.Message),
		locker:     new(sync.RWMutex),
	}
}

func (pb *rabbitPubSub) Publish(ctx context.Context, exchange string, data *pubsub.Message) error {
	go func() {
		defer common.AppRecover()
		bdata, err := data.Marshal()
		if err != nil {
			panic(err)
		}
		//if err := pb.channel.ExchangeDeclarePassive(
		//	exchange, // name
		//	"fanout", // type
		//	true,     // durable
		//	false,    // auto-deleted
		//	false,    // internal
		//	false,    // no-wait
		//	nil,      // arguments
		//); err != nil {
		//	log.Printf("Exchange %s already exists\n", exchange)
		//} else {
		//
		//	log.Println("Create exchange: ", exchange)
		//}

		err = pb.channel.ExchangeDeclare(
			"logs",   // name
			"fanout", // type
			true,     // durable
			false,    // auto-deleted
			false,    // internal
			false,    // no-wait
			nil,      // arguments
		)
		failOnError(err, "Failed to declare an exchange")

		err = pb.channel.PublishWithContext(
			ctx,
			exchange,
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bdata,
			})
		if err != nil {
			panic(err)
		}
		log.Println("New event published:", data.String(), "with data", data.Data)
	}()
	return nil
}

func (pb *rabbitPubSub) Subscribe(ctx context.Context, exchange string) (<-chan *pubsub.Message, func()) {

	q, err := pb.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = pb.channel.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := pb.channel.Consume(
		"",    // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to register a consumer")
	ch := make(chan *pubsub.Message)

	go func() {
		for {
			d := <-msgs
			var mess pubsub.Message
			mess.Unmarshal(d.Body)
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

package main

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"h5travelotobackend/component/pubsub"
	rabbitpubsub "h5travelotobackend/component/pubsub/rabbitmq"
	hotelmodel "h5travelotobackend/module/hotels/model"
	"log"
	"time"
)

func main() {
	rabbitConn, err := amqp.Dial("amqp://h5traveloto:6g9rMvbxKws56ND@localhost:5672/")
	if err != nil {
		log.Fatal("Fail to connect rabbitMQ! ", err)
	}
	defer rabbitConn.Close()
	ch, err := rabbitConn.Channel()
	if err != nil {
		log.Fatal("Fail to open channel! ", err)
	}

	pb := rabbitpubsub.NewRabbitPubSub(ch)

	msgs, _ := pb.Subscribe(context.Background(), "topic1")
	b, _ := pb.Subscribe(context.Background(), "topic1")
	go func() {
		for i := 0; i < 10; i++ {
			mess := pubsub.NewMessage(hotelmodel.Hotel{
				Name: "Hello",
			})
			mess.SetChannel("topic1")
			time.Sleep(2 * time.Second)
			pb.Publish(context.Background(), "topic1", mess)
		}
	}()

	go func() {
		for {
			d := <-msgs
			fmt.Println("a:", d.Data)
		}

	}()

	go func() {
		for {
			mess := <-b
			fmt.Println("b:", mess.Data)
		}
	}()

	time.Sleep(20 * time.Second)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"h5travelotobackend/component/pubsub"
	rabbitpubsub "h5travelotobackend/component/pubsub/rabbitmq"
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
	//b, _ := pb.Subscribe(context.Background(), "topic1")
	go func() {
		for i := 0; i < 5; i++ {
			mess := pubsub.NewMessage(1)
			mess.SetChannel("topic1")
			time.Sleep(2 * time.Second)
			pb.Publish(context.Background(), "topic1", mess)
		}
	}()

	go func() {
		for {
			d := <-msgs
			var hotel interface{}
			json.Unmarshal(d.Data, &hotel)
			fmt.Printf("Name: %T\n", hotel)
		}

	}()

	//go func() {
	//	for {
	//		mess := <-b
	//		fmt.Println("b:", mess.Data)
	//	}
	//}()

	time.Sleep(20 * time.Second)
}

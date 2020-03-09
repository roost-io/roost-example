package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ZB-io/zbio/client"
)

var (
	zbCli *client.Client
)

func init() {

	cfg := client.Config{

		Name:            "TestProducer",
		ServiceEndPoint: "zbio-service:50002",
	}
	cli, err := client.New(cfg)

	if err != nil {
		fmt.Println("Failed to get client, err=", err)
	} else {
		zbCli = cli
	}
}

func CreateTopics() {
	var topic string
	var topics []string

	for i := 1; i < 5; i++ {
		topic = fmt.Sprintf("test-topic-%d", i)
		log.Printf("Topic Name: %v", topic)
		topics = append(topics, topic)
	}

	ok, err := zbCli.CreateTopics(topics, "", 3, 1, 100000)
	if !ok {
		fmt.Errorf("Unable to create Topics due to error: %v", err)
	}
}

func SendNewMessage() {
	messages := []client.Message{
		client.Message{
			TopicName:     "test-topic-1",
			HintPartition: "",
			Data:          []byte("Message number 0"),
		},
	}

	for i := 0; i < 1000; i++ {
		messages[0].Data = []byte(fmt.Sprintf("Message number %d ", i))
		response, err := zbCli.NewMessage(messages)
		if err != nil {
			log.Fatalf("NewMessage failed with error: %v", err)
		}
		log.Printf("NewMessage Response: %v", response)

		time.Sleep(1 * time.Second)
	}
}

func startProducer() {
	CreateTopics()
	SendNewMessage()
}

func main() {

	startProducer()

}

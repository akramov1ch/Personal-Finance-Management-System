package main

import (
	"encoding/json"
	"log"
	notifactionpb "notifaction-service/protos/notifaction"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)


func main() {
    broker := "localhost:9092"
    topic := "notifications"

    p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
    if err != nil {
        log.Fatalf("Failed to create producer: %s", err)
    }
    defer p.Close()

	emailMsg := &notifactionpb.NotificationReq{
		UserEmail: "azamatgayratov771@gmail.com",
        Subject:   "Test Subject",
        Message:   "This is a test email body.",
	}

    msgBytes, err := json.Marshal(emailMsg)
    if err != nil {
        log.Fatalf("Failed to marshal email message: %s", err)
    }

    deliveryChan := make(chan kafka.Event)
    err = p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value:          msgBytes,
    }, deliveryChan)
    if err != nil {
        log.Fatalf("Failed to produce message: %s", err)
    }

    e := <-deliveryChan
    m := e.(*kafka.Message)
    if m.TopicPartition.Error != nil {
        log.Fatalf("Delivery failed: %v", m.TopicPartition.Error)
    } else {
        log.Printf("Delivered message to %v", m.TopicPartition)
    }
}

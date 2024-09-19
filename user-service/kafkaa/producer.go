package kafkaa

import (
	"encoding/json"
	"log"
	"user-service/protos/notifactionpb"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProduceRegistrationEmail(email, sub, mes string) error {
	broker := "localhost:9092"
	topic := "notifications"
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		return err
	}
	defer p.Close()

	emailMsg := &notifactionpb.NotificationReq{
		UserEmail: email,
		Subject:   sub,
		Message:   mes,
	}

	msgBytes, err := json.Marshal(emailMsg)
	if err != nil {
		log.Printf("Failed to marshal email message: %+v\n", err)
		return err
	}

	deliveryChan := make(chan kafka.Event)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgBytes,
	}, deliveryChan)
	if err != nil {
		return err
	}

	log.Println("kafkaaaagaaaa kirdiii")
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("xatolik topicpationdami...%+v", err)
		return err
	} else {
		log.Printf("Delivered message to %v", m.TopicPartition)
	}

	log.Println("kafkaaaagaaaa chiqdi")
	return nil
}

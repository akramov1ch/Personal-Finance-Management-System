package kafka

import (
	"log"
	"notifaction-service/internal/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConnKafka(cfg config.Config) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaServer.Http.Host + ":" + cfg.KafkaServer.Http.Port,
		"group.id":          cfg.KafkaServer.Group,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Kafka consumerni yaratishda xatolik: %v", err)
	}
	err = consumer.Subscribe(cfg.KafkaServer.Topic, nil)
	if err != nil {
		log.Fatalf("Kafka mavzusiga obuna bo'lishda xatolik: %v", err)
	}
	
	return consumer
}

package ckafka

import (
	"os"

	"github.com/segmentio/kafka-go"
)

var KafkaWriter *kafka.Writer

func InitKafka() {
	kafkaBroker := os.Getenv("KAFKA_BROKER")

	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

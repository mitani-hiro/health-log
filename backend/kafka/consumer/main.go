package main

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

const (
	brokerAddress = "kafka.kafka.svc.cluster.local:9092"
	topic         = "test-topic"
)

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("receive message: %s\n", string(msg.Value))
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_0_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	group, err := sarama.NewConsumerGroup([]string{brokerAddress}, "health-log-consumer-group", config)
	if err != nil {
		log.Fatalf("NewConsumerGroup error: %v", err)
	}
	defer group.Close()

	log.Println("Start Consumer")

	ctx := context.Background()
	handler := consumerGroupHandler{}

	for {
		if err := group.Consume(ctx, []string{topic}, handler); err != nil {
			log.Printf("Consume error: %v", err)
		}
	}
}

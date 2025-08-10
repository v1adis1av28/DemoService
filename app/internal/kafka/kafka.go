package kafka

import (
	"context"
	"demo/internal/models"
	"encoding/json"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

type KafkaInfo struct {
	Topic          string
	BrokkerAddress string
	GroupId        string
}

func NewKafka(ki *KafkaInfo) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{ki.BrokkerAddress},
		Topic:   ki.Topic,
		GroupID: ki.GroupId,
	})
	defer reader.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readMessages(ctx, reader)
}

func readMessages(ctx context.Context, r *kafka.Reader) {
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}
		var order models.Order
		err = json.Unmarshal(msg.Value, &order)
		fmt.Printf("Message get:\n")
		fmt.Println(order)

		if err := r.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("Failed to commit message: %v", err)
		}
	}
}

package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/LavaJover/shvark-profile-service/internal/usecase"
	"github.com/segmentio/kafka-go"
)

type UserCreatedEvent struct {
	ID			string `json:"id"`
	Login		string `json:"login"`
	Username	string `json:"username"`
	Password	string `json:"password"`
}

func StartConsumer(broker string, topic string, uc *usecase.ProfileUsecase) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "profile-service",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("error reading message:", err)
			continue
		}

		var event UserCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Println("invalid message:", err)
			continue
		}

		_ = uc.CreateProfile(context.Background(), event.Username)
	}
}
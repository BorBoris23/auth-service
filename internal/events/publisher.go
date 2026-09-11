package events

import (
	"context"
	"encoding/json"
	"strconv"

	"auth-service/internal/kafka"
)

type Message struct {
	Event string `json:"event"`
	Data  any    `json:"data"`
}

type Publisher struct {
	producer *kafka.Producer
}

func NewPublisher(producer *kafka.Producer) *Publisher {
	return &Publisher{
		producer: producer,
	}
}

func (p *Publisher) PublishUserCreated(
	ctx context.Context,
	event UserCreatedEvent,
) error {
	message := Message{
		Event: event.Name(),
		Data:  event,
	}

	eventJSON, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.producer.Publish(
		ctx,
		strconv.Itoa(event.UserID),
		eventJSON,
	)
}

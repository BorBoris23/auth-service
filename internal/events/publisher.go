package events

import (
	"context"
	"encoding/json"
	"strconv"

	"auth-service/internal/kafka"
)

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
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.producer.Publish(
		ctx,
		strconv.Itoa(event.UserID),
		eventJSON,
	)
}

package events

import (
	"context"
	"log/slog"
)

type UserCreatedListener struct {
	publisher *Publisher
}

func NewUserCreatedListener(publisher *Publisher) *UserCreatedListener {
	return &UserCreatedListener{
		publisher: publisher,
	}
}

func (l *UserCreatedListener) Handle(
	ctx context.Context,
	event Event,
) error {
	userCreatedEvent, ok := event.(*UserCreatedEvent)
	if !ok {
		return nil
	}

	err := l.publisher.PublishUserCreated(
		ctx,
		*userCreatedEvent,
	)
	if err != nil {
		slog.Error(
			"failed to publish user.created event",
			"error", err,
		)
		return err
	}

	slog.Info(
		"user.created event published",
		"user_id", userCreatedEvent.UserID,
	)

	return nil
}

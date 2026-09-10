package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(
	broker string,
	topic string,
) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &Producer{
		writer: writer,
	}
}

func (p *Producer) Publish(
	ctx context.Context,
	key string,
	value []byte,
) error {
	return p.writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: value,
		},
	)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

// KafkaProducer abstracts the kafka-go Writer
type KafkaProducer interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}

type KafkaPublisher struct {
	Producer KafkaProducer
}

func NewKafkaPublisher(producer KafkaProducer) *KafkaPublisher {
	return &KafkaPublisher{Producer: producer}
}

func (p *KafkaPublisher) Publish(
	ctx context.Context,
	topic string,
	key string,
	event interface{},
) error {
	marshaledEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}

	e := Event{
		IdempotencyKey: uuid.NewString(),
		Payload:        marshaledEvent,
		CreatedAt:      time.Now().UTC(),
	}

	data, err := json.Marshal(e)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: data,
	}

	return p.Producer.WriteMessages(ctx, msg)
}

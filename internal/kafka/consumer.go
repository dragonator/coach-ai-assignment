package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// KafkaReader abstracts the kafka-go Reader.
type KafkaReader interface {
	FetchMessage(ctx context.Context) (kafka.Message, error)
	CommitMessages(ctx context.Context, msgs ...kafka.Message) error
	Config() kafka.ReaderConfig
	Close() error
}

// EventHandler defines a function to handle consumed events.
type EventHandler func(ctx context.Context, event Event) error

type KafkaConsumer struct {
	reader  KafkaReader
	handler EventHandler
}

// NewKafkaConsumer creates a new KafkaConsumer instance.
func NewKafkaConsumer(reader KafkaReader, handler EventHandler) *KafkaConsumer {
	return &KafkaConsumer{
		reader:  reader,
		handler: handler,
	}
}

// Start begins consuming messages from Kafka.
func (c *KafkaConsumer) Start(ctx context.Context) error {
	for {
		log.Printf(`Waiting for messages on topic "%s" ...`, c.reader.Config().Topic)

		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			// If context is done, exit gracefully
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return nil
			}
			return err
		}

		var event Event
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			return fmt.Errorf("unmarshal event: %w", err)
		}

		if err := c.handler(ctx, event); err != nil {
			return fmt.Errorf("handling event: %w", err)
		}

		// Commit the message after successful handling
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			return fmt.Errorf("commit message offset: %w", err)
		}
	}
}

// Close closes the underlying Kafka reader
func (c *KafkaConsumer) Close() error {
	return c.reader.Close()
}

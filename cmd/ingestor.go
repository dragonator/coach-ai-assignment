package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"

	"github.com/dragonator/coach-ai-assignment/internal/config"
	kafkainternal "github.com/dragonator/coach-ai-assignment/internal/kafka"
)

const transactionsTopic = "transactions"

func ingestorCommand() *cobra.Command {
	ingestor := &cobra.Command{
		Use:   "ingestor",
		Short: "Starts the ingestor service",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := config.LoadConfig()
			if err != nil {
				log.Printf("Failed to load config: %v\n", err)
				os.Exit(1)
			}

			writer := kafka.NewWriter(kafka.WriterConfig{
				Brokers:  config.GetKafkaBrokers(),
				Balancer: &kafka.Hash{}, // ensure messages with the same key go to the same partition
			})
			defer writer.Close()

			publisher := kafkainternal.NewKafkaPublisher(writer)

			// Example payload
			payloads := []map[string]interface{}{
				{
					"user_id":    "123",
					"amount":     99.99,
					"created_at": time.Now().UTC(),
				},
				{
					"user_id":    "456",
					"amount":     49.99,
					"created_at": time.Now().UTC(),
				},
				{
					"user_id":    "789",
					"amount":     19.99,
					"created_at": time.Now().UTC(),
				},
			}

			// Publish the example event
			for {
				for i, payload := range payloads {
					err = publisher.Publish(context.Background(), transactionsTopic, payload["user_id"].(string), payload)
					if err != nil {
						fmt.Printf("Failed to publish event: %v\n", err)
						os.Exit(1)
					}

					time.Sleep(2 * time.Second) // Simulate some delay between messages
					fmt.Printf("Example event %d published successfully\n", i)
				}

				time.Sleep(5 * time.Second) // Wait before publishing the next batch
			}

		},
	}

	return ingestor
}

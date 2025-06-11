package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"

	"github.com/dragonator/coach-ai-assignment/internal/config"
	"github.com/dragonator/coach-ai-assignment/internal/consumer"
	kafkainternal "github.com/dragonator/coach-ai-assignment/internal/kafka"
)

func consumerCommand() *cobra.Command {
	var topic string

	consumer := &cobra.Command{
		Use:   "consumer",
		Short: "Starts a consumer instance to consume messages from a Kafka topic",
		Run: func(cmd *cobra.Command, args []string) {
			// Load configuration
			config, err := config.LoadConfig()
			if err != nil {
				fmt.Printf("Failed to load config: %v\n", err)
				os.Exit(1)
			}

			// Create a new Kafka reader
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:  config.GetKafkaBrokers(),
				Topic:    topic,
				GroupID:  config.GetKafkaConsumerGroupID(),
				MinBytes: 10e3, // 10KB
				MaxBytes: 10e6, // 10MB
			})
			defer reader.Close()

			// Initialize Prometheus registry and pusher
			prometheusPusher := push.
				New(config.GetPrometheusPushGatewayURL(), "consumer").
				Gatherer(prometheus.DefaultGatherer)

			// Initialize the consumer service with metrics decorator
			consumerService := consumer.NewService()
			consumerWithMetrics := consumer.NewMetricsDecorator(
				config,
				prometheusPusher,
				consumerService,
				topic,
			)

			// Initialize and start KafkaConsumer
			kafkaConsumer := kafkainternal.NewKafkaConsumer(reader, consumerWithMetrics.ProcessEvent)
			if err := kafkaConsumer.Start(context.Background()); err != nil {
				fmt.Printf("Consumer error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	consumer.Flags().StringVar(&topic, "topic", "", "Kafka topic to consume from")
	consumer.MarkFlagRequired("topic")

	return consumer
}

package cmd

import (
	"log"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"

	"github.com/dragonator/coach-ai-assignment/internal/config"
	"github.com/dragonator/coach-ai-assignment/internal/ingestor"
	kafkainternal "github.com/dragonator/coach-ai-assignment/internal/kafka"
	"github.com/dragonator/coach-ai-assignment/provider/client"
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
			providerClient := client.NewTransactionsClient("http://localhost:8080") // Mock server URL

			ingestorService := ingestor.NewService(publisher, providerClient)
			_ = ingestorService.IngestTransactions(cmd.Context())

			// time.Sleep(2 * time.Second) // Simulate some delay between messages
			// fmt.Printf("Example event %d published successfully\n", i)
			// }

			// time.Sleep(5 * time.Second) // Wait before publishing the next batch
			// }

		},
	}

	return ingestor
}

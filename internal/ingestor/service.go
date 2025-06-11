package ingestor

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dragonator/coach-ai-assignment/internal/events"
	"github.com/dragonator/coach-ai-assignment/provider/client"
)

type Publisher interface {
	Publish(
		ctx context.Context,
		topic string,
		key string,
		event interface{},
	) error
}

type ProviderClient interface {
	GetTransactions() ([]client.Transaction, error)
}

type Service struct {
	publisher      Publisher
	providerClient ProviderClient
}

func NewService(publisher Publisher, providerClient ProviderClient) *Service {
	return &Service{
		publisher:      publisher,
		providerClient: providerClient,
	}
}
func (s *Service) IngestTransactions(ctx context.Context) error {
	for {
		transactions, err := s.providerClient.GetTransactions()
		if err != nil {
			return err
		}

		log.Printf("Fetched %d transactions from provider\n", len(transactions))

		for _, transaction := range transactions {
			if err := validateTransaction(transaction); err != nil {
				continue
			}

			transactionEvent, err := createTransactionEvent(transaction)
			if err != nil {
				continue
			}

			if err := s.publisher.Publish(ctx, "transactions", transaction.UserID, transactionEvent); err != nil {
				return err
			}
		}

		time.Sleep(30 * time.Second) // Wait before fetching new transactions
	}

	return nil
}

func validateTransaction(transaction client.Transaction) error {
	if transaction.ID == "" ||
		transaction.UserID == "" ||
		transaction.Amount <= 0 ||
		(transaction.Type != "credit" && transaction.Type != "debit") ||
		transaction.Timestamp <= 0 {
		return errors.New("invalid transaction data")
	}
	return nil
}

func createTransactionEvent(transaction client.Transaction) (*events.TransactionEvent, error) {
	timestamp, err := time.Parse(time.RFC3339, time.Unix(transaction.Timestamp, 0).Format(time.RFC3339))
	if err != nil {
		return nil, errors.New("invalid transaction data")
	}

	return &events.TransactionEvent{
		ID:        transaction.ID,
		UserID:    transaction.UserID,
		Amount:    transaction.Amount,
		Type:      transaction.Type,
		Timestamp: timestamp,
	}, nil
}

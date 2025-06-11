package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/dragonator/coach-ai-assignment/internal/events"
	kafkainternal "github.com/dragonator/coach-ai-assignment/internal/kafka"
	"github.com/dragonator/coach-ai-assignment/internal/storage/model"
	"github.com/shopspring/decimal"
)

type UserStore interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

type TransactionStore interface {
	Upsert(ctx context.Context, transaction *model.Transaction) error
}

type Service struct {
	userStore        UserStore
	transactionStore TransactionStore
}

func NewService(userStore UserStore, transactionStore TransactionStore) *Service {
	return &Service{
		userStore:        userStore,
		transactionStore: transactionStore,
	}
}

func (s *Service) ProcessEvent(ctx context.Context, event kafkainternal.Event) error {
	log.Println("Processing event:", event)

	e := new(events.TransactionEvent)
	if err := json.Unmarshal(event.Payload, e); err != nil {
		log.Printf("Failed to unmarshal event payload: %v", err)
		return err
	}

	// check user exists or create new user
	user := model.NewUser()

	amountDecimal := decimal.NewFromFloat(e.Amount)
	transaction := model.NewTransaction(e)

	switch e.Type {
	case "debit":
		user.Balance = user.Balance.Add(amountDecimal)
	case "credit":
		user.Balance = user.Balance.Sub(amountDecimal)
	default:
		log.Printf("Unknown transaction type: %s", e.Type)
		return errors.New("unknown transaction type: " + e.Type)
	}

	if err := s.userStore.Update(ctx, user); err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}

	if err := s.transactionStore.Upsert(ctx, transaction); err != nil {
		log.Printf("Failed to upsert transaction: %v", err)
		return err
	}

	return nil
}

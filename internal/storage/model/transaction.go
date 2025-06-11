package model

import (
	"time"

	"github.com/dragonator/coach-ai-assignment/internal/events"
	"github.com/shopspring/decimal"
)

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"
)

// type A struct {
// 	value string
// }

// var (
// 	TransactionTypeDebit  = A{"debit"}
// 	TransactionTypeCredit = A{"credit"}
// )

type Transaction struct {
	ID        string          `bun:"column:id,pk,type:text"`
	UserID    string          `bun:"column:user_id,type:text"`
	Amount    decimal.Decimal `bun:"column:amount,type:numeric"`
	Type      TransactionType `bun:"column:type,type:transaction_type"`
	CreatedAt time.Time       `bun:"column:created_at,type:timestamp"`
	UpdatedAt time.Time       `bun:"column:updated_at,type:timestamp"`
}

func NewTransaction(trEvent *events.TransactionEvent) *Transaction {
	return &Transaction{
		ID:        trEvent.ID,
		UserID:    trEvent.UserID,
		Amount:    decimal.NewFromFloat(trEvent.Amount),
		Type:      TransactionType(trEvent.Type),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

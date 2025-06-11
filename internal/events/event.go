package events

import "time"

type TransactionEvent struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"` // "credit" or "debit"
	Timestamp time.Time `json:"timestamp"`
}

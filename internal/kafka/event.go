package kafka

import "time"

type Event struct {
	IdempotencyKey string      `json:"idempotency_key"`
	Payload        interface{} `json:"payload"`
	CreatedAt      time.Time   `json:"created_at"`
}

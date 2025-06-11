package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	ID        string          `bun:"column:id,pk,type:text"`
	Balance   decimal.Decimal `bun:"column:balance,type:numeric"`
	CreatedAt time.Time       `bun:"column:created_at,type:timestamp"`
	UpdatedAt time.Time       `bun:"column:updated_at,type:timestamp"`
}

func NewUser() *User {
	return &User{
		Balance:   decimal.Zero,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

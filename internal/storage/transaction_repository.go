package storage

import (
	"context"

	"github.com/dragonator/coach-ai-assignment/internal/storage/model"
	"github.com/uptrace/bun"
)

type TransactionRepository struct {
	db *bun.DB
}

func NewTransactionRepository(db *bun.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) Upsert(ctx context.Context, transaction *model.Transaction) error {
	_, err := r.db.NewInsert().Model(transaction).On("CONFLICT (id) DO UPDATE").
		Set("amount = EXCLUDED.amount").
		Set("updated_at = EXCLUDED.updated_at").
		Exec(ctx)
	return err
}

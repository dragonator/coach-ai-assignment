package consumer

import (
	"context"
	"log"

	kafkainternal "github.com/dragonator/coach-ai-assignment/internal/kafka"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ProcessEvent(_ context.Context, event kafkainternal.Event) error {
	log.Println("Processing event:", event)

	return nil
}

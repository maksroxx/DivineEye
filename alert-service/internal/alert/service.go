package alert

import (
	"context"

	"github.com/google/uuid"
	"github.com/maksroxx/DivineEye/alert-service/internal/kafka"
	"github.com/maksroxx/DivineEye/alert-service/internal/logger"
	"github.com/maksroxx/DivineEye/alert-service/internal/models"
	"github.com/maksroxx/DivineEye/alert-service/internal/repository"
	"go.uber.org/zap"
)

type Servicer interface {
	Create(ctx context.Context, userId, coin string, price float64) (string, error)
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, userId string) ([]*models.Alert, error)
}

type service struct {
	repo     repository.AlertRepository
	producer kafka.Producerer
	log      logger.Logger
}

func NewService(repo repository.AlertRepository, producer kafka.Producerer, log logger.Logger) Servicer {
	return &service{repo: repo, producer: producer, log: log}
}

func (s *service) Create(ctx context.Context, userId, coin string, price float64) (string, error) {
	id := uuid.NewString()
	a := &models.Alert{
		ID:     id,
		UserID: userId,
		Coin:   coin,
		Price:  price,
	}

	if err := s.repo.Create(ctx, a); err != nil {
		s.log.Error("failed to create alert", zap.Error(err))
		return "", err
	}
	s.log.Info("alert created", zap.String("alert_id", id))

	if err := s.producer.PublishAlertCreated(id, userId); err != nil {
		s.log.Error("failed to publish alert_created", zap.Error(err))
		return "", err
	}
	return id, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetAll(ctx context.Context, userId string) ([]*models.Alert, error) {
	return s.repo.GetAll(ctx, userId)
}

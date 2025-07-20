package repository

import (
	"context"

	"github.com/maksroxx/DivineEye/alert-service/internal/models"
)

type AlertRepository interface {
	Create(ctx context.Context, alert *models.Alert) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, userId string) ([]*models.Alert, error)
	GetById(ctx context.Context, id string) (*models.Alert, error)
}

package repository

import (
	"context"

	"github.com/maksroxx/DivineEye/notification-service/internal/models"
)

type Alerter interface {
	SaveAlert(ctx context.Context, alert models.Alert) error
	DeleteAlert(ctx context.Context, id string) error
	GetAlertsForCoin(ctx context.Context, coin string) ([]*models.Alert, error)
	GetTriggeredAlerts(ctx context.Context, coin string, price float64) ([]*models.Alert, error)
}

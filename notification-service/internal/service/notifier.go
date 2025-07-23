package service

import (
	"context"

	"github.com/maksroxx/DivineEye/notification-service/internal/fcm"
	"github.com/maksroxx/DivineEye/notification-service/internal/logger"
	"github.com/maksroxx/DivineEye/notification-service/internal/repository"
	"go.uber.org/zap"
)

type Notifier interface {
	ProcessPriceUpdate(ctx context.Context, coin string, price float64) error
}

type Notifi struct {
	Repo repository.Alerter
	Push fcm.Fcmer
	Log  logger.Logger
}

func NewNotifi(repo repository.Alerter, push fcm.Fcmer, log logger.Logger) Notifier {
	return &Notifi{
		Repo: repo,
		Push: push,
		Log:  log,
	}
}

func (n *Notifi) ProcessPriceUpdate(ctx context.Context, coin string, price float64) error {
	alerts, err := n.Repo.GetTriggeredAlerts(ctx, coin, price)
	if err != nil {
		n.Log.Error("failed to get alerts for coin", zap.String("coin", coin), zap.Error(err))
		return err
	}

	for _, alert := range alerts {
		n.Log.Info("alert", zap.String("number", alert.ID))
		err := n.Push.Send(ctx, alert.UserID, coin, price)
		if err != nil {
			n.Log.Error("failed to send notification", zap.String("user_id", alert.UserID), zap.Error(err))
		} else {
			n.Log.Info("notification sent",
				zap.String("user_id", alert.UserID),
				zap.String("coin", coin),
				zap.Float64("price", price),
			)
		}
	}
	return nil
}

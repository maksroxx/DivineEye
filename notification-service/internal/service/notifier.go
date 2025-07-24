package service

import (
	"context"

	"github.com/maksroxx/DivineEye/notification-service/internal/fcm"
	"github.com/maksroxx/DivineEye/notification-service/internal/kafka"
	"github.com/maksroxx/DivineEye/notification-service/internal/logger"
	"github.com/maksroxx/DivineEye/notification-service/internal/repository"
	"go.uber.org/zap"
)

type Notifier interface {
	ProcessPriceUpdate(ctx context.Context, coin string, price float64) error
}

type Notifi struct {
	Repo     repository.Alerter
	Push     fcm.Fcmer
	Log      logger.Logger
	Producer kafka.Producerer
}

func NewNotifi(repo repository.Alerter, push fcm.Fcmer, log logger.Logger, producer kafka.Producerer) Notifier {
	return &Notifi{
		Repo:     repo,
		Push:     push,
		Log:      log,
		Producer: producer,
	}
}

func (n *Notifi) ProcessPriceUpdate(ctx context.Context, coin string, price float64) error {
	alerts, err := n.Repo.GetTriggeredAlerts(ctx, coin, price)
	if err != nil {
		n.Log.Error("[Notifier.go] Failed to get alerts for coin", zap.String("coin", coin), zap.Error(err))
		return err
	}

	for _, alert := range alerts {
		err := n.Push.Send(ctx, alert.UserID, coin, price)
		if err != nil {
			n.Log.Error("[Notifier.go] Failed to send notification", zap.String("user_id", alert.UserID), zap.Error(err))
		} else {
			if err := n.Repo.MarkAlertTriggered(ctx, alert.ID); err != nil {
				n.Log.Error("[Notifier.go]", zap.Any("mark alert triggered", err))
			}
			_ = n.Producer.PublishAlertTriggered(alert.ID)
			n.Log.Info("[Notifier.go] Notification sent",
				zap.String("user_id", alert.UserID),
				zap.String("coin", coin),
				zap.Float64("price", price),
			)
		}
	}
	return nil
}

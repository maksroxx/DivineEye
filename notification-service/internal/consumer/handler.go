package consumer

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/notification-service/internal/logger"
	"github.com/maksroxx/DivineEye/notification-service/internal/models"
	"github.com/maksroxx/DivineEye/notification-service/internal/repository"
	"github.com/maksroxx/DivineEye/notification-service/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	Repo     repository.Alerter
	Notifier service.Notifier
	Log      logger.Logger
}

func NewHandler(repo repository.Alerter, notifier service.Notifier, log logger.Logger) *Handler {
	return &Handler{
		Repo:     repo,
		Notifier: notifier,
		Log:      log,
	}
}

func (h *Handler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h *Handler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (h *Handler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		switch msg.Topic {
		case "alerts":
			h.handleAlertEvent(sess, msg)
		case "price_updates":
			h.handlePriceEvent(sess, msg)
		default:
			h.Log.Error("[Handler.go] Received message from unknown topic", zap.String("topic", msg.Topic))
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}

func (h *Handler) handleAlertEvent(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
	type AlertPayload struct {
		AlertID   string  `json:"alert_id"`
		UserID    string  `json:"user_id"`
		Coin      string  `json:"coin"`
		Price     float64 `json:"price"`
		Direction string  `json:"direction"`
		Type      string  `json:"type"`
	}

	var payload AlertPayload
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		h.Log.Error("[Handler.go] Failed to unmarshal alert event", zap.Error(err))
		return
	}

	switch payload.Type {
	case "alert_created":
		alert := models.Alert{
			ID:        payload.AlertID,
			UserID:    payload.UserID,
			Coin:      payload.Coin,
			Direction: payload.Direction,
			Price:     payload.Price,
		}
		if err := h.Repo.SaveAlert(context.Background(), alert); err != nil {
			h.Log.Error("[Handler.go] Failed to save alert", zap.Error(err))
		} else {
			h.Log.Info("[Handler.go] Alert created", zap.String("alert_id", payload.AlertID))
		}

	case "alert_deleted":
		if err := h.Repo.DeleteAlert(context.Background(), payload.AlertID); err != nil {
			h.Log.Error("[Handler.go] Failed to delete alert", zap.Error(err))
		} else {
			h.Log.Info("[Handler.go] Alert deleted", zap.String("alert_id", payload.AlertID))
		}

	default:
		h.Log.Info("[Handler.go] Unknown alert type", zap.String("type", payload.Type))
	}
}

func (h *Handler) handlePriceEvent(sess sarama.ConsumerGroupSession, msg *sarama.ConsumerMessage) {
	var payload struct {
		Symbol string  `json:"symbol"`
		Price  float64 `json:"price"`
	}
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		h.Log.Error("[Handler.go] Failed to unmarshal price event", zap.Error(err))
		return
	}
	h.Notifier.ProcessPriceUpdate(context.Background(), payload.Symbol, payload.Price)
}

package kafka

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/maksroxx/DivineEye/alert-service/internal/repository"
)

type TriggeredHandler struct {
	Repo repository.AlertRepository
}

func NewTriggeredHandler(repo repository.AlertRepository) *TriggeredHandler {
	return &TriggeredHandler{Repo: repo}
}

func (h *TriggeredHandler) Setup(sess sarama.ConsumerGroupSession) error   { return nil }
func (h *TriggeredHandler) Cleanup(sess sarama.ConsumerGroupSession) error { return nil }

func (h *TriggeredHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var evt struct {
			AlertID string `json:"alert_id"`
			Event   string `json:"event"`
		}
		if err := json.Unmarshal(msg.Value, &evt); err != nil {
			continue
		}
		if evt.Event == "triggered" {
			_ = h.Repo.MarkAlertTriggered(context.Background(), evt.AlertID)
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}

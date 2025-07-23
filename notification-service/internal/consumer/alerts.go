package consumer

// import (
// 	"context"
// 	"encoding/json"
// 	"strconv"

// 	"github.com/IBM/sarama"
// 	"github.com/maksroxx/DivineEye/notification-service/internal/models"
// 	"github.com/maksroxx/DivineEye/notification-service/internal/repository"
// )

// type AlertsHandler struct {
// 	Repo repository.Alerter
// }

// func (h *AlertsHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
// func (h *AlertsHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// func (h *AlertsHandler) ConsumerClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
// 	for msg := range claim.Messages() {
// 		var payload map[string]string
// 		_ = json.Unmarshal(msg.Value, &payload)

// 		alertID := payload["alert_id"]
// 		userID := payload["user_id"]
// 		coin := payload["coin"]
// 		price := payload["price"]
// 		direction := payload["direction"]
// 		typ := payload["type"]

// 		switch typ {
// 		case "alert_created":
// 			p, _ := strconv.ParseFloat(price, 64)
// 			_ = h.Repo.SaveAlert(context.Background(), models.Alert{
// 				ID:        alertID,
// 				UserID:    userID,
// 				Coin:      coin,
// 				Direction: direction,
// 				Price:     p,
// 			})
// 		case "alert_deleted":
// 			_ = h.Repo.DeleteAlert(context.Background(), alertID)
// 		}
// 		sess.MarkMessage(msg, "")
// 	}
// 	return nil
// }

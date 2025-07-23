package consumer

// import (
// 	"context"
// 	"encoding/json"

// 	"github.com/IBM/sarama"
// 	"github.com/maksroxx/DivineEye/notification-service/internal/service"
// )

// type PricesHandler struct {
// 	Notifier service.Notifier
// }

// func (h *PricesHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
// func (h *PricesHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// func (h *PricesHandler) ConsumerClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
// 	for msg := range claim.Messages() {
// 		var payload struct {
// 			Symbol string  `json:"symbol"`
// 			Price  float64 `json:"price"`
// 		}
// 		if err := json.Unmarshal(msg.Value, &payload); err != nil {
// 			continue
// 		}

// 		h.Notifier.ProcessPriceUpdate(context.Background(), payload.Symbol, payload.Price)
// 		sess.MarkMessage(msg, "")
// 	}
// 	return nil
// }

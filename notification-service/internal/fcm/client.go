package fcm

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

type Fcmer interface {
	Send(ctx context.Context, userID, coin string, price float64) error
}

type Sender struct {
	client *messaging.Client
}

func NewSenderr(ctx context.Context, credentialsPath string) (*Sender, error) {
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		return nil, fmt.Errorf("failed to init firebase app: %w", err)
	}
	msgClient, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to init firebase messaging client: %w", err)
	}
	return &Sender{client: msgClient}, nil
}

func NewSender(ctx context.Context, credentialsPath string) (*Sender, error) {
	// app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credentialsPath))
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to init firebase app: %w", err)
	// }
	// msgClient, err := app.Messaging(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to init firebase messaging client: %w", err)
	// }
	return &Sender{client: nil}, nil
}

func (s *Sender) Send(ctx context.Context, userID, coin string, price float64) error {
	return nil
}

func (s *Sender) Sendd(ctx context.Context, userID, coin string, price float64) error {
	msg := &messaging.Message{
		Topic: fmt.Sprintf("%s", userID),
		Notification: &messaging.Notification{
			Title: "📈 Price Alert",
			Body:  fmt.Sprintf("%s reached %.2f", coin, price),
		},
		Data: map[string]string{
			"coin":  coin,
			"price": fmt.Sprintf("%.2f", price),
		},
	}

	_, err := s.client.Send(ctx, msg)
	if err != nil {
		return fmt.Errorf("fcm send failed: %w", err)
	}
	return nil
}

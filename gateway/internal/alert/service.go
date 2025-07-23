package alert

import (
	"context"

	pb "github.com/maksroxx/DivineEye/alert-service/proto"
	"github.com/maksroxx/DivineEye/gateway/internal/grpcclient"
)

type AlertService interface {
	Create(ctx context.Context, userID, coin, direction string, price float64) (string, error)
	Get(ctx context.Context, userID string) ([]*pb.Alert, error)
	Delete(ctx context.Context, id string) error
}

type alertService struct {
	client *grpcclient.AlertClient
}

func NewAlertService(client *grpcclient.AlertClient) AlertService {
	return &alertService{client: client}
}

func (s *alertService) Create(ctx context.Context, userID, coin, direction string, price float64) (string, error) {
	resp, err := s.client.CreateAlert(ctx, &pb.CreateAlertRequest{
		UserId:    userID,
		Coin:      coin,
		Price:     price,
		Direction: direction,
	})
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func (s *alertService) Get(ctx context.Context, userID string) ([]*pb.Alert, error) {
	resp, err := s.client.GetAlerts(ctx, &pb.GetAlertsRequest{UserId: userID})
	if err != nil {
		return nil, err
	}
	return resp.Alerts, nil
}

func (s *alertService) Delete(ctx context.Context, id string) error {
	_, err := s.client.DeleteAlert(ctx, &pb.DeleteAlertRequest{Id: id})
	return err
}

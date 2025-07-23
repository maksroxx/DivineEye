package alert

import (
	"context"

	proto "github.com/maksroxx/DivineEye/alert-service/proto"
)

type GRPCServer struct {
	proto.UnimplementedAlertServiceServer
	svc Servicer
}

func NewGRPCServer(svc Servicer) *GRPCServer {
	return &GRPCServer{svc: svc}
}

func (s *GRPCServer) CreateAlert(ctx context.Context, req *proto.CreateAlertRequest) (*proto.CreateAlertResponse, error) {
	id, err := s.svc.Create(ctx, req.UserId, req.Coin, req.Direction, req.Price)
	if err != nil {
		return nil, err
	}
	return &proto.CreateAlertResponse{Id: id}, nil
}

func (s *GRPCServer) DeleteAlert(ctx context.Context, req *proto.DeleteAlertRequest) (*proto.DeleteAlertResponse, error) {
	return &proto.DeleteAlertResponse{}, s.svc.Delete(ctx, req.Id)
}

func (s *GRPCServer) GetAlerts(ctx context.Context, req *proto.GetAlertsRequest) (*proto.GetAlertsResponse, error) {
	alerts, err := s.svc.GetAll(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var pbAlerts []*proto.Alert
	for _, a := range alerts {
		pbAlerts = append(pbAlerts, &proto.Alert{
			Id:        a.ID,
			Coin:      a.Coin,
			Direction: a.Direction,
			Price:     a.Price,
		})
	}
	return &proto.GetAlertsResponse{Alerts: pbAlerts}, nil
}

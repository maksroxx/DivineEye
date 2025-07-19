package grpcclient

import (
	pb "github.com/maksroxx/DivineEye/alert-service/proto"
	"google.golang.org/grpc"
)

type AlertClient struct {
	pb.AlertServiceClient
}

func NewAlertClient(addr string) *AlertClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &AlertClient{
		AlertServiceClient: pb.NewAlertServiceClient(conn),
	}
}

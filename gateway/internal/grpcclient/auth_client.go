package grpcclient

import (
	pb "github.com/maksroxx/DivineEye/gateway/proto-clients/auth"
	"google.golang.org/grpc"
)

type AuthClient struct {
	pb.AuthServiceClient
}

func NewAuthClient(addr string) *AuthClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &AuthClient{AuthServiceClient: pb.NewAuthServiceClient(conn)}
}

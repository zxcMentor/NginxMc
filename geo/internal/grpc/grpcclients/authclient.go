package grpcclients

import (
	"context"
	pbauth "github.com/zxcMentor/grpcproto/protos/auth/gen/go"
	"google.golang.org/grpc"
	"log"
)

type ClientAuth struct{}

func (c *ClientAuth) CallIsValid(ctx context.Context, req *pbauth.ValidRequest) (*pbauth.ValidResponse, error) {
	conn, err := grpc.Dial("auth:50052", grpc.WithInsecure())
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	client := pbauth.NewAuthServiceClient(conn)
	res, err := client.ItsValid(ctx, req)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return res, nil
}

package grpcclients

import (
	"context"
	pbuser "github.com/zxcMentor/grpcproto/protos/user/gen/go"
	"google.golang.org/grpc"
	"log"
)

type ClientUser struct{}

func NewClientUser() *ClientUser {
	return &ClientUser{}
}

func (c *ClientUser) CallCreateUser(ctx context.Context, req *pbuser.CreateUserRequest) (*pbuser.CreateUserResponse, error) {
	conn, err := grpc.Dial("user:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err connect grpc:", err)
	}

	client := pbuser.NewUserServiceClient(conn)

	res, err := client.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatal("err call grpc", err)
	}
	return res, nil
}

func (c *ClientUser) CallCheckUser(ctx context.Context, req *pbuser.CheckUserRequest) (*pbuser.CheckUserResponse, error) {
	conn, err := grpc.Dial("user1:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err connect grpc:", err)
	}

	client := pbuser.NewUserServiceClient(conn)

	res, err := client.CheckUser(context.Background(), req)
	if err != nil {
		log.Fatal("err call grpc")
	}
	return res, nil
}

func (c *ClientUser) CallProfileUser(ctx context.Context, req *pbuser.ProfileUserRequest) (*pbuser.ProfileUserResponse, error) {
	conn, err := grpc.Dial("user1:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatal("err connect grpc:", err)
	}

	client := pbuser.NewUserServiceClient(conn)

	res, err := client.ProfileUser(context.Background(), req)
	if err != nil {
		log.Fatal("err call grpc")
	}
	return res, nil
}

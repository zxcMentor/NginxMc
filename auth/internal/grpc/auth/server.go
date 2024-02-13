package auth

import (
	"auth/internal/service"
	"context"
	pbauth "github.com/zxcMentor/grpcproto/protos/auth/gen/go"
	"log"
)

type ServicerAuth interface {
	Register(email, password string) (string, error)
	Login(email, password string) (string, error)
	ItsValid(token string) (bool, error)
}

type ServiceAuth struct {
	pbauth.UnimplementedAuthServiceServer
	auths service.AuthService
}

func (s *ServiceAuth) Register(ctx context.Context, req *pbauth.RegisterRequest) (*pbauth.RegisterResponse, error) {
	mess, err := s.auths.Register(req.Email, req.Hashepassword)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	return &pbauth.RegisterResponse{Message: mess}, nil
}

func (s *ServiceAuth) Login(ctx context.Context, req *pbauth.LoginRequest) (*pbauth.LoginResponse, error) {
	token, err := s.auths.Login(req.Email, req.Password)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return &pbauth.LoginResponse{Token: token}, nil
}

func (s *ServiceAuth) ItsValid(ctx context.Context, req *pbauth.ValidRequest) (*pbauth.ValidResponse, error) {
	isvalid, err := s.auths.ItsValid(req.Token)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return &pbauth.ValidResponse{IsValid: isvalid}, nil
}

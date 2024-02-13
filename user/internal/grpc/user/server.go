package user

import (
	"context"
	"user/internal/models"
	"user/internal/service"
	pbuser "user/protos/gen/go"
)

type ServicerUser interface {
	CreateUser(email, password string) (string, error)
	CheckUser(email, password string) error
	ProfileUser(email string) (*models.UserDTO, error)
	ListUsers() (*[]models.UserDTO, error)
}

type ServiceUser struct {
	pbuser.UnimplementedUserServiceServer
	us *service.UserService
}

func NewServiceUser(usservice *service.UserService) *ServiceUser {
	return &ServiceUser{us: usservice}
}

func (s *ServiceUser) CreateUser(ctx context.Context, req *pbuser.CreateUserRequest) (*pbuser.CreateUserResponse, error) {
	message, err := s.us.CreateUser(req.Email, req.HashPassword)
	if err != nil {
		return nil, err
	}
	return &pbuser.CreateUserResponse{Message: message}, nil
}

func (s *ServiceUser) CheckUser(ctx context.Context, req *pbuser.CheckUserRequest) (*pbuser.CheckUserResponse, error) {
	err := s.us.CheckUser(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pbuser.CheckUserResponse{}, nil
}

func (s *ServiceUser) ProfileUser(ctx context.Context, req *pbuser.ProfileUserRequest) (*pbuser.ProfileUserResponse, error) {
	user, err := s.us.ProfileUser(req.Email)
	if err != nil {
		return nil, err
	}
	p := &pbuser.User{Id: user.Id, Email: user.Email}

	return &pbuser.ProfileUserResponse{User: p}, nil
}
func (s *ServiceUser) ListUsers(ctx context.Context, req *pbuser.ListUsersRequest) (*pbuser.ListUsersResponse, error) {
	users, err := s.us.ListUsers()
	if err != nil {
		return nil, err
	}
	var grpcUsers []*pbuser.User
	for _, user := range users {
		grpcUser := &pbuser.User{
			Id:    user.Id,
			Email: user.Email,
		}
		grpcUsers = append(grpcUsers, grpcUser)
	}
	return &pbuser.ListUsersResponse{Users: grpcUsers}, nil

}

package service

import (
	"auth/internal/grpc/grpcclients"
	"context"
	"github.com/golang-jwt/jwt"
	pbuser "github.com/zxcMentor/grpcproto/protos/user/gen/go"
	"log"
	"time"
)

const SecretKey = "secretkey"

type AuthService struct {
	clientUser *grpcclients.ClientUser
}

func NewAuthService(grcl *grpcclients.ClientUser) *AuthService {
	return &AuthService{clientUser: grcl}
}

func (a *AuthService) Register(email, hashepassword string) (string, error) {

	mess, err := a.clientUser.CallCreateUser(context.Background(), &pbuser.CreateUserRequest{
		Email:        email,
		HashPassword: hashepassword,
	})
	if err != nil {
		log.Fatal("err call user")
	}

	return mess.Message, nil
}

func (a *AuthService) Login(email, password string) (string, error) {

	_, err := a.clientUser.CallCheckUser(context.Background(), &pbuser.CheckUserRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Println("err: ", err)
		return "", err
	}
	token, err := GenerateToken(email)
	if err != nil {
		log.Println("err generate token:", err)
		return "", err
	}

	return token, nil
}

func (a *AuthService) ItsValid(token string) (bool, error) {

	return false, nil
}

func GenerateToken(email string) (string, error) {

	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}

	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

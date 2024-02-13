package service

import (
	"fmt"

	"log"
	"user/internal/models"
	"user/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo}
}

func (u *UserService) CreateUser(email, hashepassword string) (string, error) {
	err := u.Repo.CreateUser(email, hashepassword)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	return fmt.Sprint("user created successfully"), nil
}

func (u *UserService) CheckUser(email, password string) error {
	err := u.Repo.CheckUser(email, password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ProfileUser(email string) (*models.UserDTO, error) {
	user, err := u.Repo.ProfileUser(email)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserService) ListUsers() ([]models.UserDTO, error) {
	users, err := u.Repo.ListUsers()
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return users, nil
}

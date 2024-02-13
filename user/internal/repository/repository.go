package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"user/internal/models"
)

type UserRepository interface {
	CreateUser(email, hashepassword string) error
	CheckUser(email, password string) error
	ProfileUser(email string) (*models.UserDTO, error)
	ListUsers() (*[]models.UserDTO, error)
}

type UserRepo struct {
	Postgres *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) CreateUser(email, hashepassword string) error {
	query := `INSERT INTO users (email, hashepassword) VALUES ($1, $2)`
	result, err := u.Postgres.Exec(query, email, hashepassword)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	// Получение количества затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %v", err)
	}

	return nil

}

func (u *UserRepo) CheckUser(email, password string) error {
	var user models.User
	query := `SELECT id, email, hashepassword FROM users WHERE email = $1`
	err := u.Postgres.Get(&user, query, email)
	if err != nil {
		log.Println("err not found user")
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashePassword), []byte(password))
	if err != nil {
		log.Printf("invalid password: %v", err)
		return err
	}
	return nil
}

func (u *UserRepo) ProfileUser(email string) (*models.UserDTO, error) {
	var user models.UserDTO
	query := `SELECT id, email FROM users WHERE email = $1`
	err := u.Postgres.Get(&user, query, email)
	if err != nil {
		log.Println("err user not exist")
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) ListUsers() ([]models.UserDTO, error) {
	var users []models.UserDTO
	query := `SELECT id, email FROM users `
	err := u.Postgres.Select(&users, query)
	if err != nil {
		log.Println("err dont get users")
		return nil, err
	}

	return users, nil
}

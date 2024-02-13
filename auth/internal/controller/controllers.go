package controller

import (
	"auth/internal/service"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type HandleAuth struct {
	servAuth *service.AuthService
}

func NewHandleAuth(servAuth *service.AuthService) *HandleAuth {
	return &HandleAuth{servAuth: servAuth}
}

func (h *HandleAuth) Register(w http.ResponseWriter, r *http.Request) {
	email := "@example"
	password := "1234"

	hashepassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("err generate hashepassword")
	}

	mess, err := h.servAuth.Register(email, string(hashepassword))
	if err != nil {
		http.Error(w, "err service", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(mess))
}

func (h *HandleAuth) Login(w http.ResponseWriter, r *http.Request) {
	email := "@example"
	password := "1234"

	token, err := h.servAuth.Login(email, password)
	if err != nil {
		http.Error(w, "err register failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

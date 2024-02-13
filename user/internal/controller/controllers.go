package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"user/internal/service"
)

type HandleUser struct {
	servUser *service.UserService
}

func NewHandleUser(sUser *service.UserService) *HandleUser {
	return &HandleUser{sUser}
}

func (h *HandleUser) Profile(w http.ResponseWriter, r *http.Request) {
	email := "@example"

	user, err := h.servUser.ProfileUser(email)
	if err != nil {
		http.Error(w, "err serv", http.StatusInternalServerError)
	}
	jsData, err := json.Marshal(user)
	if err != nil {
		log.Fatal("err", err)
	}
	w.Write(jsData)
}

func (h *HandleUser) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.servUser.ListUsers()
	if err != nil {
		log.Println("err list")
	}
	jsData, err := json.Marshal(users)
	w.Write(jsData)
}

package router

import (
	"github.com/go-chi/chi"
	"user/internal/controller"
)

func StRout(cn *controller.HandleUser) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/api/profile", cn.Profile)
	r.Get("/api/list", cn.List)
	return r
}

package router

import (
	"auth/internal/controller"
	"github.com/go-chi/chi"
	"net/http"
)

func StRout(cn *controller.HandleAuth) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/register", cn.Register)
	r.Get("/api/login", cn.Login)
	r.Get("/api/os", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("auth ok"))
	})
	return r
}

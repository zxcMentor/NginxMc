package router

import (
	"geo/internal/controller"
	"github.com/go-chi/chi"
)

func StartRout(cont *controller.HandleGeo) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/address/search", cont.SearchHandle)
	r.Post("/api/address/search", cont.GeocodeHandle)

	return r
}

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"geo/internal/models"
	"geo/internal/service"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

type HandleGeo struct {
	serv  *service.GeoService
	redcl *redis.Client
}

func NewHandGeo(serviceGeo *service.GeoService) *HandleGeo {
	redcl := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return &HandleGeo{
		serv:  serviceGeo,
		redcl: redcl,
	}
}

func (h *HandleGeo) SearchHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sear run")
	req := &models.SearchRequest{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	cacheKey := fmt.Sprintf("geoSearch: %s", req.Query)
	data, err := h.redcl.Get(context.Background(), cacheKey).Result()
	if err == redis.Nil {

		address, err := h.serv.GeoSearch(req.Query)
		if err != nil {
			http.Error(w, "err Call GRPC", http.StatusInternalServerError)
			return
		}

		h.redcl.Set(context.Background(), cacheKey, address, 20*time.Second)

		w.Write(address)
	} else if err != nil {

		http.Error(w, "Cache retrieval error", http.StatusInternalServerError)
	} else {

		w.Write([]byte(data))
	}

}

func (h *HandleGeo) GeocodeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("geocode run")
	req := &models.GeocodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	cacheKey := fmt.Sprintf("geoGeocode: %s %s", req.Lat, req.Lng)
	data, err := h.redcl.Get(context.Background(), cacheKey).Result()
	if err == redis.Nil {
		// Данных нет в кеше, выполняем запрос к gRPC сервису
		address, err := h.serv.GeoCode(req.Lat, req.Lng)
		if err != nil {
			http.Error(w, "err Call GRPC", http.StatusInternalServerError)
			return
		}

		h.redcl.Set(context.Background(), cacheKey, address, 20*time.Second)

		w.Write(address)
	} else if err != nil {

		http.Error(w, "Cache retrieval error", http.StatusInternalServerError)
	} else {

		w.Write([]byte(data))
	}
}

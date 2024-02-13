package main

import (
	"geo/internal/controller"
	"geo/internal/grpc/geo"
	"geo/internal/router"
	"geo/internal/service"
	pbgeo "geo/protos/gen/go"
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {
	ns := service.NewGeoService()
	nh := controller.NewHandGeo(ns)
	r := router.StartRout(nh)
	w := sync.WaitGroup{}
	w.Add(2)
	go func(r *chi.Mux) {
		defer w.Done()
		http.ListenAndServe(":8081", r)
	}(r)

	go func() {
		defer w.Done()
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Ошибка при прослушивании порта: %v", err)
		}

		server := grpc.NewServer()
		pbgeo.RegisterGeoServiceServer(server, &geo.ServerGeo{})

		log.Println("Запуск gRPC сервера geo...")
		if err := server.Serve(listen); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()
	w.Wait()
}

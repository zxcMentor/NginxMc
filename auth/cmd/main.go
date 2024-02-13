package main

import (
	"auth/internal/controller"
	"auth/internal/grpc/auth"
	"auth/internal/grpc/grpcclients"
	"auth/internal/router"
	"auth/internal/service"
	"fmt"
	"github.com/go-chi/chi"
	pbauth "github.com/zxcMentor/grpcproto/protos/auth/gen/go"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

func main() {

	clUs := grpcclients.NewClientUser()
	as := service.NewAuthService(clUs)
	ac := controller.NewHandleAuth(as)
	r := router.StRout(ac)
	w := sync.WaitGroup{}
	w.Add(2)
	go func(r *chi.Mux) {
		fmt.Println("Запусе сервера auth")
		http.ListenAndServe(":8082", r)
		defer w.Done()
	}(r)

	go func() {
		listen, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("Ошибка при прослушивании порта: %v", err)
		}

		server := grpc.NewServer()
		pbauth.RegisterAuthServiceServer(server, &auth.ServiceAuth{})

		log.Println("Запуск gRPC сервера auth...")
		if err := server.Serve(listen); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
		defer w.Done()
	}()
	w.Wait()
}

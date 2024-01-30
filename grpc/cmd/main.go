package main

import (
	"HUGO/grpc/internal/service"
	pb "HUGO/grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGeoServiceServer(server, &service.GeoService{})

	log.Println("Запуск gRPC сервера...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

package main

import (
	"json-rpc/internal/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	geoProvider := new(service.GeoProvide)
	err := rpc.Register(geoProvider)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":4321")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}

	log.Println("Сервер запущен на порту 1234")
	http.Serve(l, nil)
}

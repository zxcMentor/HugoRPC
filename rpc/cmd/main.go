package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc/internal/rpc/geo"
)

func main() {

	geoProvider := new(geo.ServerGeo)
	err := rpc.Register(geoProvider)
	if err != nil {
		panic(err)
	}
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}

	log.Println("Сервер запущен на порту 1234")
	rpc.Accept(l)

}

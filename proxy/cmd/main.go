package main

import (
	"log"
	"net/http"
	"proxy/internal/controller"
	"proxy/internal/router"
	"proxy/internal/service"
)

func main() {

	gs := service.NewGeoService()
	gh := controller.NewGeoHand(&gs)
	r := router.StRout(gh)

	log.Println("serv 8080 started")
	http.ListenAndServe(":8080", r)
}

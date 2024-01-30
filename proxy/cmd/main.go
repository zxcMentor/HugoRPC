package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"proxy/internal/controller"
	"proxy/internal/router"
	"proxy/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading env:", err)
	}
	gs := service.NewGeoService()
	gh := controller.NewGeoHand(&gs)
	r := router.StRout(gh)

	log.Println("proxy serv started on ports :8080")
	http.ListenAndServe(":8080", r)
}

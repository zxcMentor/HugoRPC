package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	_ "proxy/docs"
	"proxy/internal/controller"
	"proxy/internal/router"
)

// @title Proxy Service API
// @version 1.0
// @description This is the API documentation for the Proxy Service.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("err loading env:", err)
	}

	gh := controller.NewGeoHand()
	r := router.StRout(gh)

	log.Println("proxy serv started on ports :8080")
	http.ListenAndServe(":8080", r)
}

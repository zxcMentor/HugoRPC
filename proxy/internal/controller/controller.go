package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"proxy/internal/models"
	"proxy/internal/service"
)

type GeoHandler interface {
	SearchHandler(w http.ResponseWriter, r http.Request)
}

type GeoHandle struct {
	GeoServ service.GeoService
}

func NewGeoHand(GeoServ *service.GeoService) *GeoHandle {

	return &GeoHandle{GeoServ: *GeoServ}
}

func (h *GeoHandle) SearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handle run")
	var req models.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	address, err := h.GeoServ.GeoSearch(req.Query)
	if err != nil {
		log.Fatal("err:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(address)
	if err != nil {
		log.Fatal("err:", err)
	}

}

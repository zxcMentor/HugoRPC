package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"proxy/internal/models"
	"proxy/internal/rpc/rpcClient"
	"proxy/internal/service"
)

type GeoHandler interface {
	SearchHandler(w http.ResponseWriter, r *http.Request)
	GeocodeHandler(w http.ResponseWriter, r *http.Request)
}

type GeoHandle struct {
	GeoServ service.GeoService
}

func NewGeoHand(GeoServ *service.GeoService) *GeoHandle {

	return &GeoHandle{GeoServ: *GeoServ}
}

// @Summary Search handler
// @Description Handles search requests
// @ID search-handler
// @Accept json
// @Produce json
// @Param   query     body    models.SearchRequest true        "Search query"
// @Success 200 {object} models.Address
// @Failure 400 {string} string
// @Router /api/address/search [post]
func (h *GeoHandle) SearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handle run")
	var req models.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	protocol := os.Getenv("RPC_PROTOCOL")
	addresses := []*models.Address{}
	switch protocol {
	case "rpc":
		rpcFactory := rpcClient.NewClientRpcFactory()
		address, err := rpcFactory.CreateClientAndCallSearch(req.Query)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		addresses = address

	case "json-rpc":
		jsonrpcFactory := rpcClient.NewJsonRpcClientFactory()
		address, err := jsonrpcFactory.CreateClientAndCallSearch(req.Query)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		addresses = address

	case "grpc":
		grpcFactory := rpcClient.NewGrpcClientFactory()
		address, err := grpcFactory.CreateClientAndCallSearch(req.Query)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		addresses = address

	default:
		log.Println("unknown protocol")
		http.Error(w, "Unsupported Protocol", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(addresses)
	if err != nil {
		log.Println("err:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// @Summary Geocode handler
// @Description Handles geocode requests
// @ID geocode-handler
// @Accept json
// @Produce json
// @Param   lat     body    float64 true        "Latitude"
// @Param   lng     body    float64 true        "Longitude"
// @Success 200 {object} models.Address
// @Failure 400 {string} string
// @Router /api/address/geocode [post]
func (h *GeoHandle) GeocodeHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err decoding :", err)
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}
	address, err := h.GeoServ.GeoCode(req.Lat, req.Lng)
	if err != nil {
		log.Fatal("dont get address:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(address)
	if err != nil {
		log.Println("err encode :", err)
		return
	}
}

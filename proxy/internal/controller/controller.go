package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proxy/internal/models"
	"proxy/internal/rpc/rpcClient"
)

type GeoHandler interface {
	SearchHandler(w http.ResponseWriter, r *http.Request)
	GeocodeHandler(w http.ResponseWriter, r *http.Request)
}

type GeoHandle struct {
}

func NewGeoHand() *GeoHandle {

	return &GeoHandle{}
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
	fmt.Println("SearchHandler run")
	var req models.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	protocol := "rpc"

	var result []byte
	switch protocol {
	case "rpc":
		rpcFactory := rpcClient.NewClientRpcFactory()
		address, err := rpcFactory.CreateClientAndCallSearch(req.Query)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = address

	case "json-rpc":
		jsonrpcFactory := rpcClient.NewJsonRpcClientFactory()
		address, err := jsonrpcFactory.CreateClientAndCallSearch(req.Query)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = address

	case "grpc":
		grpcFactory := rpcClient.NewGrpcClientFactory()
		address, err := grpcFactory.CreateClientAndCallSearch(req.Query)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = address

	default:
		log.Println("unknown protocol")
		http.Error(w, "Unsupported Protocol", http.StatusNotImplemented)
		return
	}
	fmt.Println(string(result))
	w.Write(result)

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
	fmt.Println("GeocodeHandler run")
	var req *models.GeocodeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	protocol := "rpc"

	var result []byte
	switch protocol {
	case "rpc":
		rpcFactory := rpcClient.NewClientRpcFactory()
		address, err := rpcFactory.CreateClientAndCallGeocode(req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = address

	case "json-rpc":
		jsonrpcFactory := rpcClient.NewJsonRpcClientFactory()
		address, err := jsonrpcFactory.CreateClientAndCallGeocode(req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = address

	case "grpc":
		grpcFactory := rpcClient.NewGrpcClientFactory()
		address, err := grpcFactory.CreateClientAndCallGeocode(req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = address

	default:
		log.Println("unknown protocol")
		http.Error(w, "Unsupported Protocol", http.StatusNotImplemented)
		return
	}
	fmt.Println(string(result))
	w.Write(result)

}

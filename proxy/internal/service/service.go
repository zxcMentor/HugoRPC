package service

import (
	"fmt"
	"log"
	"os"
	"proxy/internal/grpcclient"
	"proxy/internal/models"
	"proxy/internal/rpcclient"
)

type GeoServicer interface {
	GeoSearch(input string) ([]*models.Address, error)
	GeoCode(lat, lng string) ([]*models.Address, error)
}

type GeoService struct {
}

func NewGeoService() GeoService {
	return GeoService{}
}

func (g *GeoService) GeoSearch(input string) ([]*models.Address, error) {
	protos := os.Getenv("RPC_PROTOCOL")

	switch protos {
	case "rpc":
		address, err := rpcclient.CreateRpcClient(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, err
	case "json-rpc":

	case "grpc":
		address, err := grpcclient.CreateGRPCClient(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, nil

	}
	return nil, fmt.Errorf("unknown rpc protocol: %s", protos)
}

func (g *GeoService) GeoCode(lat, lng string) ([]*models.Address, error) {
	return nil, nil
}

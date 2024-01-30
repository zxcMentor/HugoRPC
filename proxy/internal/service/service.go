package service

import (
	"fmt"
	"log"
	"os"
	"proxy/internal/models"
	"proxy/internal/rpc/grpcClient"
	"proxy/internal/rpc/rpcClient"
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
	protocol := os.Getenv("RPC_PROTOCOL")

	switch protocol {
	case "rpc":
		address, err := rpcClient.ConnectAndCallRpc(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, err
	case "json-rpc":

	case "grpc":
		address, err := grpcClient.ConnectAndCallGRPC(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, nil

	}
	return nil, fmt.Errorf("unknown rpc protocol: %s", protocol)
}

func (g *GeoService) GeoCode(lat, lng string) ([]*models.Address, error) {
	return nil, nil
}

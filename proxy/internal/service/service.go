package service

import (
	"fmt"
	"log"
	"os"
	"proxy/internal/models"
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
		rpcFactory := rpcClient.NewClientRpcFactory()
		address, err := rpcFactory.CreateClientAndCallSearch(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, nil

	case "json-rpc":
		jsonrpcFactory := rpcClient.NewJsonRpcClientFactory()
		address, err := jsonrpcFactory.CreateClientAndCallSearch(input)
		if err != nil {
			log.Fatal("err:", err)
			return nil, err
		}
		return address, nil
	case "grpc":
		grpcFactory := rpcClient.NewGrpcClientFactory()
		address, err := grpcFactory.CreateClientAndCallSearch(input)
		if err != nil {
			log.Fatal("err:", err)
		}
		return address, nil

	}
	return nil, fmt.Errorf("unknown rpc protocol: %s", protocol)
}

func (g *GeoService) GeoCode(lat, lng string) ([]*models.Address, error) {

	return nil, nil
}

package rpcClient

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net/rpc"
	"proxy/internal/models"
	"proxy/internal/rpc/rpcClient/proto"
)

type ClientFactoryRpc interface {
	CreateClientAndCallSearch(input string) ([]byte, error)
	CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error)
}

type ClientGrpcFactory struct{}

func NewGrpcClientFactory() *ClientGrpcFactory {
	return &ClientGrpcFactory{}
}

func (f *ClientGrpcFactory) CreateClientAndCallSearch(input string) ([]byte, error) {
	conn, err := grpc.Dial("grpc:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto.NewGeoServiceClient(conn)

	req := &proto.GeoAddressRequest{Input: input}
	res, err := client.GeoAddressSearch(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
		return nil, err
	}
	var result []byte

	result = res.Data

	return result, nil
}

func (f *ClientGrpcFactory) CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error) {
	conn, err := grpc.Dial("grpc:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto.NewGeoServiceClient(conn)

	req := &proto.GeocodeRequest{
		Lat: inp.Lat,
		Lon: inp.Lng,
	}
	res, err := client.GeoAddressGeocode(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
		return nil, err
	}
	var result []byte

	result = res.Data

	return result, nil
}

type ClientJsonRpcFactory struct{}

func NewJsonRpcClientFactory() *ClientJsonRpcFactory {
	return &ClientJsonRpcFactory{}
}

func (f *ClientJsonRpcFactory) CreateClientAndCallSearch(input string) ([]byte, error) {
	client, err := rpc.DialHTTP("tcp", "json-rpc:4321")
	if err != nil {
		log.Fatal(err)
	}

	var address []byte
	err = client.Call("ServerGeo.SearchGeoAddress", input, &address)
	if err != nil {
		log.Fatal(err)
	}

	return address, err
}

func (f *ClientJsonRpcFactory) CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error) {
	client, err := rpc.DialHTTP("tcp", "json-rpc:4321")
	if err != nil {
		log.Fatal(err)
	}

	var address []byte
	err = client.Call("ServerGeo.GeocodeAddress", inp, &address)
	if err != nil {
		log.Fatal(err)
	}

	return address, err
}

type ClientRpcFactory struct {
}

func NewClientRpcFactory() *ClientRpcFactory {
	return &ClientRpcFactory{}
}

func (f *ClientRpcFactory) CreateClientAndCallSearch(input string) ([]byte, error) {
	client, err := rpc.Dial("tcp", "rpc:1234")
	if err != nil {
		log.Fatal("err dial:", err)
		return nil, err
	}
	var address []byte
	err = client.Call("ServerGeo.SearchGeoAddress", input, &address)
	if err != nil {
		log.Fatal("err call:", err)
		return nil, err
	}
	return address, nil
}

func (f *ClientRpcFactory) CreateClientAndCallGeocode(inp *models.GeocodeRequest) ([]byte, error) {
	client, err := rpc.Dial("tcp", "rpc:1234")
	if err != nil {
		log.Fatal("err dial:", err)
		return nil, err
	}
	var address []byte
	err = client.Call("ServerGeo.GeocodeAddress", inp, &address)
	if err != nil {
		log.Fatal("err call:", err)
		return nil, err
	}
	return address, nil
}

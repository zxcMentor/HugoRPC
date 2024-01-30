package rpcClient

import (
	"log"
	"net/rpc"
	"proxy/internal/models"
)

func ConnectAndCallRpc(input string) ([]*models.Address, error) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("err:", err)
		return nil, err
	}
	var address []*models.Address
	err = client.Call("GeoProvider.SearchGeoAddress", input, address)
	if err != nil {
		log.Fatal("err:", err)
		return nil, err
	}
	return address, nil
}

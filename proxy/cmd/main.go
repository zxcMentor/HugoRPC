package main

import (
	"log"
	"net/http"
	"os"
	_ "proxy/docs"
	"proxy/internal/controller"
	"proxy/internal/router"
	"proxy/internal/rpc/rpcClient"
)

// @title Proxy Service API
// @version 1.0
// @description This is the API documentation for the Proxy Service.
// @host localhost:8080
// @BasePath /
func main() {
	protocol := os.Getenv("RPC_PROTOCOL")
	var rpccl rpcClient.ClientFactoryRpc

	switch protocol {
	case "rpc":
		rpcFactory := rpcClient.NewClientRpcFactory()
		rpccl = rpcFactory

	case "json-rpc":
		jsonrpcFactory := rpcClient.NewJsonRpcClientFactory()

		rpccl = jsonrpcFactory

	case "grpc":
		grpcFactory := rpcClient.NewGrpcClientFactory()
		rpccl = grpcFactory

	default:
		log.Println("unknown protocol")
		return
	}
	gh := controller.NewGeoHand(rpccl)
	r := router.StRout(gh)

	log.Println("proxy serv started on ports :8080")
	http.ListenAndServe(":8080", r)
}

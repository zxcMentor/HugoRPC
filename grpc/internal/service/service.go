package service

import (
	pb "HUGO/grpc/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GeoService struct {
	pb.UnimplementedGeoServiceServer
}

func (g *GeoService) GeoAddressSearch(ctx context.Context, request *pb.GeoAddressRequest) (*pb.GeoAddressResponse, error) {

	jsd, err := json.Marshal(request.Input)
	if err != nil {
		log.Fatal("err marshaling:", err)
	}

	req, err := http.NewRequest("POST", "http://suggestions.dadata.ru/suggestions/api/4_1/rs/suggest/address", bytes.NewBuffer(jsd))
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+"5086f9aa3d01c20cab4d1477df59cb0f1ab75497:01c3fde0996a6e08e1d51d5203c57cdb211739b2")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}

	var sugg pb.GeoAddressResponse
	err = json.Unmarshal(body, &sugg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &pb.GeoAddressResponse{Addresses: sugg.Addresses}, nil
}

func (g *GeoService) GeoAddressGeocode(ctx context.Context, req *pb.GeocodeRequest) (*pb.GeocodeResponse, error) {
	return nil, nil
}

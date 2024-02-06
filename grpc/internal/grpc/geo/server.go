package geo

import (
	"HUGO/grpc/internal/service"
	pb "HUGO/grpc/proto"
	"context"
	"fmt"
)

type Geo interface {
	GeoSearch(input string) ([]byte, error)
	GeoCode(lat, lng string) ([]byte, error)
}

type ServiceGeo struct {
	pb.UnimplementedGeoServiceServer
	geo service.GeoService
}

func (g *ServiceGeo) GeoAddressSearch(ctx context.Context, request *pb.GeoAddressRequest) (*pb.GeoAddressResponse, error) {

	address, err := g.geo.GeoSearch(request.Input)
	if err != nil {
		return nil, fmt.Errorf("err get address:%v", err)
	}

	return &pb.GeoAddressResponse{Data: address}, nil
}

func (g *ServiceGeo) GeoAddressGeocode(ctx context.Context, req *pb.GeocodeRequest) (*pb.GeocodeResponse, error) {
	address, err := g.geo.GeoCode(req.Lat, req.Lon)
	if err != nil {
		return nil, fmt.Errorf("err get address:%v", err)
	}

	return &pb.GeocodeResponse{Data: address}, nil
}

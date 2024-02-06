package geo

import (
	"log"
	"rpc/internal/models"
	"rpc/internal/service"
)

type ServerGeo struct {
	geo service.GeoProvide
}

func (g *ServerGeo) SearchGeoAddress(args string, reply *[]byte) error {
	address, err := g.geo.SearchGeoAddress(args)
	if err != nil {
		log.Fatal("err call rpc:", err)
	}
	*reply = address
	return nil
}

func (g *ServerGeo) GeocodeAddress(inp *models.GeocodeRequest, reply *[]byte) error {
	address, err := g.geo.GeocodeAddress(inp)
	if err != nil {
		log.Fatal("err call rpc:", err)
	}
	*reply = address
	return nil
}

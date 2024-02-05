package geo

import (
	"log"
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
	reply = &address
	return nil
}

func (g *ServerGeo) GeocodeAddress() error {
	//TODO implement me
	return nil
}

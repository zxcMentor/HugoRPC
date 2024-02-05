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
	CreateClientAndCallSearch(input string) ([]*models.Address, error)
	CreateClientAndCallGeocode(lat, lon string) ([]*models.Address, error)
}

type ClientGrpcFactory struct{}

func NewGrpcClientFactory() *ClientGrpcFactory {
	return &ClientGrpcFactory{}
}

func (f *ClientGrpcFactory) CreateClientAndCallSearch(input string) ([]*models.Address, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
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

	var addresses []*models.Address
	for _, addr := range res.Addresses {
		address := &models.Address{
			PostalCode:           addr.PostalCode,
			Country:              addr.Country,
			CountryISOCode:       addr.CountryIsoCode,
			FederalDistrict:      addr.FederalDistrict,
			RegionFIASID:         addr.RegionFiasId,
			RegionKLADRID:        addr.RegionKladrId,
			RegionISOCode:        addr.RegionIsoCode,
			RegionWithType:       addr.RegionWithType,
			RegionType:           addr.RegionType,
			RegionTypeFull:       addr.RegionTypeFull,
			Region:               addr.Region,
			AreaFIASID:           addr.AreaFiasId,
			AreaKLADRID:          addr.AreaKladrId,
			AreaWithType:         addr.AreaWithType,
			AreaType:             addr.AreaType,
			AreaTypeFull:         addr.AreaTypeFull,
			Area:                 addr.Area,
			CityFIASID:           addr.CityFiasId,
			CityKLADRID:          addr.CityKladrId,
			CityWithType:         addr.CityWithType,
			CityType:             addr.CityType,
			CityTypeFull:         addr.CityTypeFull,
			City:                 addr.City,
			CityArea:             addr.CityArea,
			CityDistrictFIASID:   addr.CityDistrictFiasId,
			CityDistrictKLADRID:  addr.CityDistrictKladrId,
			CityDistrictWithType: addr.CityDistrictWithType,
			CityDistrictType:     addr.CityDistrictType,
			CityDistrictTypeFull: addr.CityDistrictTypeFull,
			CityDistrict:         addr.CityDistrict,
			StreetFIASID:         addr.StreetFiasId,
			StreetKLADRID:        addr.StreetKladrId,
			StreetWithType:       addr.StreetWithType,
			StreetType:           addr.StreetType,
			StreetTypeFull:       addr.StreetTypeFull,
			Street:               addr.Street,
			SteadFIASID:          addr.SteadFiasId,
			SteadCadnum:          addr.SteadCadnum,
			SteadType:            addr.SteadType,
			SteadTypeFull:        addr.SteadTypeFull,
			Stead:                addr.Stead,
			HouseFIASID:          addr.HouseFiasId,
			HouseKLADRID:         addr.HouseKladrId,
			HouseCadnum:          addr.HouseCadnum,
			HouseType:            addr.HouseType,
			HouseTypeFull:        addr.HouseTypeFull,
			House:                addr.House,
			BlockType:            addr.BlockType,
			BlockTypeFull:        addr.BlockTypeFull,
			Block:                addr.Block,
			Entrance:             addr.Entrance,
			Floor:                addr.Floor,
			FlatFIASID:           addr.FlatFiasId,
			FlatCadnum:           addr.FlatCadnum,
			FlatType:             addr.FlatType,
			FlatTypeFull:         addr.FlatTypeFull,
			Flat:                 addr.Flat,
			FlatArea:             addr.FlatArea,
			SquareMeterPrice:     addr.SquareMeterPrice,
			FlatPrice:            addr.FlatPrice,
			PostalBox:            addr.PostalBox,
			FIASID:               addr.FiasId,
			FIASCadastreNumber:   addr.FiasCadastreNumber,
			FIASLevel:            addr.FiasLevel,
			FIASActualityState:   addr.FiasActualityState,
			KLADRID:              addr.KladrId,
			GeonameID:            addr.GeonameId,
			CapitalMarker:        addr.CapitalMarker,
			OKATO:                addr.Okato,
			OKTMO:                addr.Oktmo,
			TaxOffice:            addr.TaxOffice,
			TaxOfficeLegal:       addr.TaxOfficeLegal,
			Timezone:             addr.Timezone,
			GeoLat:               addr.GeoLat,
			GeoLon:               addr.GeoLon,
			BeltwayHit:           addr.BeltwayHit,
			BeltwayDistance:      addr.BeltwayDistance,
			Metro:                addr.Metro,
			Divisions:            addr.Divisions,
			QCGeo:                addr.QcGeo,
			QCComplete:           addr.QcComplete,
			QCHouse:              addr.QcHouse,
			HistoryValues:        addr.HistoryValues,
			UnparsedParts:        addr.UnparsedParts,
			Source:               addr.Source,
			QC:                   addr.Qc,
		}
		addresses = append(addresses, address)
	}

	log.Printf("Ответ от сервера: %s", res.Addresses)
	return addresses, nil
}

func (f *ClientGrpcFactory) CreateClientAndCallGeocode(lat, lon string) ([]*models.Address, error) {
	//TODO implement me
	return nil, nil
}

type ClientJsonRpcFactory struct{}

func NewJsonRpcClientFactory() *ClientJsonRpcFactory {
	return &ClientJsonRpcFactory{}
}

func (f *ClientJsonRpcFactory) CreateClientAndCallSearch(input string) ([]*models.Address, error) {
	client, err := rpc.DialHTTP("tcp", "json-rpc:4321")
	if err != nil {
		log.Fatal(err)
	}

	var address []*models.Address
	err = client.Call("ServerGeo.SearchGeoAddress", input, &address)
	if err != nil {
		log.Fatal(err)
	}

	return address, err
}

func (f *ClientJsonRpcFactory) CreateClientAndCallGeocode(lat, lon string) ([]*models.Address, error) {
	//TODO implement me
	return nil, nil
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

func (f *ClientRpcFactory) CreateClientAndCallGeocode(lat, lon string) ([]*models.Address, error) {
	//TODO implement me
	return nil, nil
}

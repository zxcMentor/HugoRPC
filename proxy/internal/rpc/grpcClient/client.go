package grpcClient

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"proxy/internal/models"
	"proxy/internal/rpc/grpcClient/proto"
)

func ConnectAndCallGRPC(input string) ([]*models.Address, error) {
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

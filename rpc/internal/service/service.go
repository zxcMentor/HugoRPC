package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rpc/internal/models"
	"strings"
)

type GeoProvide struct {
}

func (g *GeoProvide) SearchGeoAddress(input string) ([]byte, error) {
	var data = strings.NewReader(fmt.Sprintf("[ \"%s\" ]", input))

	req, err := http.NewRequest("POST", "https://cleaner.dadata.ru/api/v1/clean/address", data)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+"24133541982a4f8baa0497ac37c7661de6598b13")
	req.Header.Set("X-Secret", "bbff5cda452ec7ecbf4eea2f3c4e97538e599b46")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("dadata err request:", err)
	}
	defer resp.Body.Close()
	log.Println("dadata statuscode - ", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}

	var adrs models.AddressSearch
	err = json.Unmarshal(body, &adrs)
	if err != nil {
		log.Println("err unmarshal ")
	}

	var ress models.AddressSearchElement
	for _, s := range adrs {
		ress = models.AddressSearchElement{
			Source:               "",
			Result:               s.Result,
			PostalCode:           "",
			Country:              "",
			CountryISOCode:       "",
			FederalDistrict:      "",
			RegionFiasID:         "",
			RegionKladrID:        "",
			RegionISOCode:        "",
			RegionWithType:       "",
			RegionType:           "",
			RegionTypeFull:       "",
			Region:               "",
			AreaFiasID:           nil,
			AreaKladrID:          nil,
			AreaWithType:         nil,
			AreaType:             nil,
			AreaTypeFull:         nil,
			Area:                 nil,
			CityFiasID:           nil,
			CityKladrID:          nil,
			CityWithType:         nil,
			CityType:             nil,
			CityTypeFull:         nil,
			City:                 nil,
			CityArea:             "",
			CityDistrictFiasID:   nil,
			CityDistrictKladrID:  nil,
			CityDistrictWithType: "",
			CityDistrictType:     "",
			CityDistrictTypeFull: "",
			CityDistrict:         "",
			SettlementFiasID:     nil,
			SettlementKladrID:    nil,
			SettlementWithType:   nil,
			SettlementType:       nil,
			SettlementTypeFull:   nil,
			Settlement:           nil,
			StreetFiasID:         "",
			StreetKladrID:        "",
			StreetWithType:       "",
			StreetType:           "",
			StreetTypeFull:       "",
			Street:               "",
			SteadFiasID:          nil,
			SteadKladrID:         nil,
			SteadCadnum:          nil,
			SteadType:            nil,
			SteadTypeFull:        nil,
			Stead:                nil,
			HouseFiasID:          "",
			HouseKladrID:         "",
			HouseCadnum:          "",
			HouseType:            "",
			HouseTypeFull:        "",
			House:                "",
			BlockType:            nil,
			BlockTypeFull:        nil,
			Block:                nil,
			Entrance:             nil,
			Floor:                nil,
			FlatFiasID:           "",
			FlatCadnum:           "",
			FlatType:             "",
			FlatTypeFull:         "",
			Flat:                 "",
			FlatArea:             "",
			SquareMeterPrice:     "",
			FlatPrice:            "",
			PostalBox:            nil,
			FiasID:               "",
			FiasCode:             "",
			FiasLevel:            "",
			FiasActualityState:   "",
			KladrID:              "",
			CapitalMarker:        "",
			Okato:                "",
			Oktmo:                "",
			TaxOffice:            "",
			TaxOfficeLegal:       "",
			Timezone:             "",
			GeoLat:               s.GeoLat,
			GeoLon:               s.GeoLon,
			BeltwayHit:           "",
			BeltwayDistance:      nil,
			QcGeo:                0,
			QcComplete:           0,
			QcHouse:              0,
			Qc:                   0,
			UnparsedParts:        nil,
			Metro:                nil,
		}
	}

	address, err := json.Marshal(ress)
	if err != nil {
		log.Println("err marshal :", err)
	}
	return address, nil
}

func (g *GeoProvide) GeocodeAddress(inp *models.GeocodeRequest) ([]byte, error) {
	var data = strings.NewReader(fmt.Sprintf("{ \"lat\":%s, \"lon\":%s }", inp.Lat, inp.Lng)) //(`{ "lat": 55.878, "lon": 37.653 }`)
	req, err := http.NewRequest("POST", "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token 24133541982a4f8baa0497ac37c7661de6598b13")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err read response body", err)
	}
	var adrs models.GeocodeResponse
	err = json.Unmarshal(body, &adrs)
	if err != nil {
		log.Println("err unmarshal ", err)
	}

	var ress models.Suggestion
	for _, s := range adrs.Suggestions {
		ress = models.Suggestion{
			Value:             s.Value,
			UnrestrictedValue: s.UnrestrictedValue,
			Data:              s.Data,
		}

	}

	address, err := json.Marshal(ress)
	if err != nil {
		log.Println("err marshal :", err)
	}
	return address, nil
}

package service

import (
	"encoding/json"
	"fmt"
	"io"
	"json-rpc/internal/models"
	"log"
	"net/http"
	"strings"
)

type GeoProvide struct {
}

func (g *GeoProvide) SearchGeoAddress(args string) ([]byte, error) {
	var data = strings.NewReader(fmt.Sprintf("[ \"%s\" ]", args))

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

	var pr models.ResponseAddress

	for _, s := range adrs {

		prs := []models.AddressSearchEl{{
			Result: s.Result,
			GeoLat: s.GeoLat,
			GeoLon: s.GeoLon,
		},
		}

		pr.Addresses = prs
	}

	addresses, err := json.Marshal(pr)
	if err != nil {
		log.Println("err marshal :", err)
	}

	return addresses, nil
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

	var pr models.ResponseGeocode

	for _, v := range adrs.Suggestions {
		as := []models.AddressSearchEl{
			{
				Result: v.Value,
				GeoLat: v.Data.GeoLat,
				GeoLon: v.Data.GeoLon,
			},
		}
		ps := models.ResponseGeocode{
			Value:             v.Value,
			UnrestrictedValue: v.UnrestrictedValue,
			Data:              as,
		}

		pr = ps
	}

	address, err := json.Marshal(pr)
	if err != nil {
		log.Println("err marshal :", err)
	}

	return address, nil
}

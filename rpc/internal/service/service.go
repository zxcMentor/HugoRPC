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

	var adr []*models.AddressSearchElement
	err = json.Unmarshal(body, &adr)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}
	for _, f := range adr {
		fmt.Println(f.Result)
	}
	var adrs models.AddressSearch
	res := make([]models.Address, 0)
	for _, s := range adrs {
		res = append(res, models.Address(s.Result))
	}
	fmt.Println(res)

	return json.Marshal(res)
}

func (g *GeoProvide) GeocodeAddress() error {
	//TODO implement me
	l := &AA{BB: CC{
		GG:  "",
		GSD: "",
	}}
	return nil
}

type AA struct {
	BB CC
}
type CC struct {
	GG  string
	GSD string
}

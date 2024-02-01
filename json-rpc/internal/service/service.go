package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"json-rpc/internal/models"
	"log"
	"net/http"
)

type GeoProvide struct {
}

func (g *GeoProvide) SearchGeoAddress(args string) (*[]*models.Address, error) {
	sreq := models.SearchRequest{Query: args}
	jsd, _ := json.Marshal(sreq)

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

	var sugg models.SearchResponse
	err = json.Unmarshal(body, &sugg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}
	reply := &[]*models.Address{}
	for _, s := range sugg.Suggestions {
		*reply = append(*reply, s.Data)
	}

	return reply, nil
}

func (g *GeoProvide) GeocodeAddress() error {
	//TODO implement me
	return nil
}

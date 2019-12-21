package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/johanavril/airvisual"
)

func main() {
	client := airvisual.New(
		"API KEY",
		airvisual.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
	)

	countries, err := client.Countries()
	if err != nil {
		log.Panicf("Failed to get supported Countries: %v\n", err)
	}
	if countries.Status != "success" {
		log.Panicf("Countries API reject the request: %s\n", countries.Status)
	}

	countryName := countries.Data[36].Country
	states, err := client.States(countryName)
	if err != nil {
		log.Panicf("Failed to get supported States: %v\n", err)
	}
	if states.Status != "success" {
		log.Panicf("States API reject the request: %s\n", states.Status)
	}

	stateName := states.Data[2].State
	cities, err := client.Cities(stateName, countryName)
	if err != nil {
		log.Panicf("Failed to get supported Cities: %v\n", err)
	}
	if cities.Status != "success" {
		log.Panicf("Cities API reject the request: %s\n", cities.Status)
	}

	cityName := cities.Data[0].City
	city, err := client.City(cityName, stateName, countryName)
	if err != nil {
		log.Panicf("Failed to get supported City: %v\n", err)
	}
	if city.Status != "success" {
		log.Panicf("City API reject the request: %s\n", city.Status)
	}

	c, err := json.MarshalIndent(city, "", "\t")
	if err != nil {
		log.Panicf("Failed to marshal City: %v\n", err)
	}

	fmt.Printf("%s", string(c))
}
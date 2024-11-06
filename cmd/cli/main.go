package main

import (
	"fmt"
	"log"

	"github.com/ruanzerah/fw/internal"
	"github.com/ruanzerah/fw/internal/api"
)

func main() {
	req, err := api.SendRequest(15, -7.8825, 40.0817)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.Daily.MaxTemperature)

	str := internal.SWeatherToString(req.Daily.WeatherCode)
	fmt.Println(str)
	fmt.Println(req.Daily.MinTemperature)
	fmt.Println("_______________________________________________________")
	fmt.Printf("daily: %v , current: %v ", req.Daily, req.Current)
}

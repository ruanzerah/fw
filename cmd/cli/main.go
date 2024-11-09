package main

import (
	"flag"
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ruanzerah/fw/internal"
	"github.com/ruanzerah/fw/internal/api"
)

func main() {
	// TODO: find a better way to search weather with locations (lat , long)
	loc := flag.String("Location", "Xique-Xique", "Location to get coordinates")
	flag.Parse()
	lat, long, err := api.GetGeoCoordinates(*loc)
	if err != nil {
		log.Fatal(err)
	}
	req, err := api.SendRequest(lat, long)
	if err != nil {
		log.Fatal(err)
	}
	p := internal.CreateCurrentPanel(req.Current)

	n := internal.CreateDailyPanel(req.Daily)
	fmt.Print(p)
	fmt.Print("")
	fmt.Println(n)
}

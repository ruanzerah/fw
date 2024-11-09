package main

import (
	"fmt"
	"log"

	"github.com/ruanzerah/fw/internal"
	"github.com/ruanzerah/fw/internal/api"
)

func main() {
	// TODO: find a better way to search weather with locations (lat , long)
	req, err := api.SendRequest(5, 7)
	if err != nil {
		log.Fatal(err)
	}
	p := internal.CreateCurrentPanel(req.Current)

	n := internal.CreateDailyPanel(req.Daily)
	fmt.Print(p)
	fmt.Print("")
	fmt.Println(n)
}

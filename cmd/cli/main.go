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
	p := internal.CreateCurrentPanel(req.Current)

	n := internal.CreateDailyPanel(req.Daily)
	fmt.Print(p)
	fmt.Print("")
	fmt.Println(n)
}

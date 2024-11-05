package main

import (
	"fmt"

	"github.com/ruanzerah/fw/internal/api"
)

func main() {
	req := api.SendRequest(-7.8825, 40.0817)

	fmt.Printf("daily: %v , current: %v ", req.Daily, req.Current)
}

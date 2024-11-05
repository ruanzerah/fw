package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ruanzerah/fw/internal"
)

func SendRequest(lat float32, long float32) *internal.WheaterResponse {
	latStr := strconv.FormatFloat(float64(lat), 'f', -1, 32)
	longStr := strconv.FormatFloat(float64(long), 'f', -1, 32)
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + latStr + "&longitude=" + longStr + "&current=temperature_2m,weather_code&daily=uv_index_max&timezone=auto&forecast_days=1"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var r internal.WheaterResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}

	return &r
}

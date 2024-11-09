package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ruanzerah/fw/internal"
)

func SendRequest(lat, long float64) (*internal.WeatherResponse, error) {
	latStr := strconv.FormatFloat(lat, 'f', -1, 32)
	longStr := strconv.FormatFloat(long, 'f', -1, 32)
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + latStr + "&longitude=" + longStr + "&current=temperature_2m,apparent_temperature,weather_code&daily=weather_code,temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,uv_index_max&timezone=auto"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var r internal.WeatherResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}

	return &r, nil
}

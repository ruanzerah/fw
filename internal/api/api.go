package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ruanzerah/fw/internal"
)

func SendRequest(days int, lat, long float32) (*internal.WeatherResponse, error) {
	if days > 16 {
		return nil, errors.New("forecast_days must be lower than 16. Ex: (1, 3, 7, 16)")
	}
	latStr := strconv.FormatFloat(float64(lat), 'f', -1, 32)
	longStr := strconv.FormatFloat(float64(long), 'f', -1, 32)
	daysStr := strconv.FormatInt(int64(days), 10)
	url := "https://api.open-meteo.com/v1/forecast?latitude=" + latStr + "&longitude=" + longStr + "&current=temperature_2m,apparent_temperature,weather_code&daily=weather_code,temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,uv_index_max&timezone=auto&forecast_days" + daysStr
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

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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

func GetGeoCoordinates(loc string) (float64, float64, error) {
	apiKey := os.Getenv("API_KEY")
	locQuery := url.QueryEscape(loc)
	url := "https://api.geoapify.com/v1/geocode/search?text=" + locQuery + "&format=json&apiKey=" + apiKey
	res, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}

	defer res.Body.Close()
	var r internal.GeoLoc

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}

	if len(r.Results) == 0 {
		return 0, 0, fmt.Errorf("no results found for location: %s", loc)
	}

	lat := r.Results[0].Lat
	lon := r.Results[0].Lon
	return lat, lon, nil
}

package internal

type WeatherResponse struct {
	Current Current `json:"current"`
	Daily   Daily   `json:"daily"`
}

type (
	Current struct {
		WeatherCode         WeatherCode `json:"weather_code"`
		Temperature         float32     `json:"temperature_2m"`
		ApparentTemperature float32     `json:"apparent_temperature"`
	}
)

type Daily struct {
	Date           []string      `json:"time"`
	UvIndex        []float32     `json:"uv_index_max"`
	MaxTemperature []float32     `json:"temperature_2m_max"`
	MinTemperature []float32     `json:"temperature_2m_min"`
	WeatherCode    []WeatherCode `json:"weather_code"`
}

type WeatherCode int

const (
	ClearSky     WeatherCode = 0
	MainlyClear  WeatherCode = 1
	PartlyCloudy WeatherCode = 2
	Overcast     WeatherCode = 3

	Fog               WeatherCode = 45
	DepositingRimeFog WeatherCode = 48

	DrizzleLight    WeatherCode = 51
	DrizzleModerate WeatherCode = 53
	DrizzleDense    WeatherCode = 55

	FreezingDrizzleLight WeatherCode = 56
	FreezingDrizzleDense WeatherCode = 57

	RainSlight   WeatherCode = 61
	RainModerate WeatherCode = 63
	RainHeavy    WeatherCode = 65

	FreezingRainLight WeatherCode = 66
	FreezingRainHeavy WeatherCode = 67

	SnowFallSlight   WeatherCode = 71
	SnowFallModerate WeatherCode = 73
	SnowFallHeavy    WeatherCode = 75

	SnowGrains WeatherCode = 77

	RainShowersSlight   WeatherCode = 80
	RainShowersModerate WeatherCode = 81
	RainShowersViolent  WeatherCode = 82

	SnowShowersSlight           WeatherCode = 85
	SnowShowersHeavy            WeatherCode = 86
	ThunderstormSlightModerated WeatherCode = 95
	ThunderstormSlight          WeatherCode = 96
	ThunderstormHeavyHail       WeatherCode = 99
)

var weatherDescriptions = map[WeatherCode]string{
	ClearSky:                    "Clear sky",
	MainlyClear:                 "Mainly clear",
	PartlyCloudy:                "Partly cloudy",
	Overcast:                    "Overcast",
	Fog:                         "Fog",
	DepositingRimeFog:           "Depositing rime fog",
	DrizzleLight:                "Drizzle: Light intensity",
	DrizzleModerate:             "Drizzle: Moderate intensity",
	DrizzleDense:                "Drizzle: Dense intensity",
	FreezingDrizzleLight:        "Freezing Drizzle: Light intensity",
	FreezingDrizzleDense:        "Freezing Drizzle: Dense intensity",
	RainSlight:                  "Rain: Slight intensity",
	RainModerate:                "Rain: Moderate intensity",
	RainHeavy:                   "Rain: Heavy intensity",
	FreezingRainLight:           "Freezing Rain: Light intensity",
	FreezingRainHeavy:           "Freezing Rain: Heavy intensity",
	SnowFallSlight:              "Snow fall: Slight intensity",
	SnowFallModerate:            "Snow fall: Moderate intensity",
	SnowFallHeavy:               "Snow fall: Heavy intensity",
	SnowGrains:                  "Snow grains",
	RainShowersSlight:           "Rain showers: Slight intensity",
	RainShowersModerate:         "Rain showers: Moderate intensity",
	RainShowersViolent:          "Rain showers: Violent intensity",
	SnowShowersSlight:           "Snow showers: Slight intensity",
	SnowShowersHeavy:            "Snow showers: Heavy intensity",
	ThunderstormSlightModerated: "Thunderstorm: Slight or moderate",
	ThunderstormSlight:          "Thunderstorm with slight and heavy hail",
	ThunderstormHeavyHail:       "Thunderstorm with slight and heavy hail",
}

func WeatherToString(wc WeatherCode) string {
	if desc, exists := weatherDescriptions[wc]; exists {
		return desc
	}
	return "Unknown weather code"
}

func SWeatherToString(wc []WeatherCode) []string {
	result := make([]string, len(wc))
	for i, code := range wc {
		result[i] = WeatherToString(WeatherCode(code))
	}
	return result
}

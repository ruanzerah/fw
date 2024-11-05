package internal

type WheaterResponse struct {
	Current Current `json:"current"`
	Daily   Daily   `json:"daily"`
}
type (
	WeatherCode int
	Current     struct {
		Temperature float32     `json:"temperature_2m"`
		WheaterCode WeatherCode `json:"wheater_code"`
	}
)

type Daily struct {
	UvIndex []float32 `json:"uv_index_max"`
}

const (
	ClearSky WeatherCode = 0

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

	SnowShowersSlight WeatherCode = 85
	SnowShowersHeavy  WeatherCode = 86
)

func (wc WeatherCode) String() string {
	switch wc {
	case ClearSky:
		return "Clear sky"
	case MainlyClear:
		return "Mainly clear"
	case PartlyCloudy:
		return "Partly cloudy"
	case Overcast:
		return "Overcast"
	case Fog:
		return "Fog"
	case DepositingRimeFog:
		return "Depositing rime fog"
	case DrizzleLight:
		return "Drizzle: Light intensity"
	case DrizzleModerate:
		return "Drizzle: Moderate intensity"
	case DrizzleDense:
		return "Drizzle: Dense intensity"
	case FreezingDrizzleLight:
		return "Freezing Drizzle: Light intensity"
	case FreezingDrizzleDense:
		return "Freezing Drizzle: Dense intensity"
	case RainSlight:
		return "Rain: Slight intensity"
	case RainModerate:
		return "Rain: Moderate intensity"
	case RainHeavy:
		return "Rain: Heavy intensity"
	case FreezingRainLight:
		return "Freezing Rain: Light intensity"
	case FreezingRainHeavy:
		return "Freezing Rain: Heavy intensity"
	case SnowFallSlight:
		return "Snow fall: Slight intensity"
	case SnowFallModerate:
		return "Snow fall: Moderate intensity"
	case SnowFallHeavy:
		return "Snow fall: Heavy intensity"
	case SnowGrains:
		return "Snow grains"
	case RainShowersSlight:
		return "Rain showers: Slight intensity"
	case RainShowersModerate:
		return "Rain showers: Moderate intensity"
	case RainShowersViolent:
		return "Rain showers: Violent intensity"
	case SnowShowersSlight:
		return "Snow showers: Slight intensity"
	case SnowShowersHeavy:
		return "Snow showers: Heavy intensity"
	default:
		return "Unknown weather code"
	}
}

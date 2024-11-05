package internal

type WheaterResponse struct {
	Current Current `json:"current"`
	Daily   Daily   `json:"daily"`
}

type Current struct {
	Temperature float32 `json:"temperature_2m"`
	WheaterCode int     `json:"wheater_code"`
}

type Daily struct {
	UvIndex []float32 `json:"uv_index_max"`
}

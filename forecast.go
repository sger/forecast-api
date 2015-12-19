package forecast

const URL_FORCAST_API = "https://api.forecast.io/forecast"

type Forecast struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Offset    float64 `json:"offset"`
}

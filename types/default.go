package types

type CurrentData struct {
	Interval            int     `json:"interval"`
	RelativeHumidity2m  int     `json:"relative_humidity_2m"`
	Temperature2m       float64 `json:"temperature_2m"`
	Time                string  `json:"time"`
	WindSpeed10m        float64 `json:"wind_speed_10m"`
}

type CurrentUnits struct {
	Interval           string `json:"interval"`
	RelativeHumidity2m string `json:"relative_humidity_2m"`
	Temperature2m      string `json:"temperature_2m"`
	Time               string `json:"time"`
	WindSpeed10m       string `json:"wind_speed_10m"`
}

type WeatherData struct {
	Current        CurrentData  `json:"current"`
	CurrentUnits   CurrentUnits `json:"current_units"`
	Elevation      float64      `json:"elevation"`
	GenerationTime float64      `json:"generationtime_ms"`
	Latitude       float64      `json:"latitude"`
	Longitude      float64      `json:"longitude"`
	Timezone       string       `json:"timezone"`
	TimezoneAbbrev string       `json:"timezone_abbreviation"`
	UtcOffset      int          `json:"utc_offset_seconds"`
}
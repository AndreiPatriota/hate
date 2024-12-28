package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"hate/controllers"
	"hate/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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


func main() {
	models.InitDb()

	e := echo.New()
	e.Use(middleware.Static("/static"))

	e.GET("/", controllers.ShowIndex)
	e.GET("/home", controllers.ShowHome)
	e.GET("/produtos", controllers.ShowProdutos)
	e.GET("/sobre", controllers.ShowSobre)
	e.GET("/tempo", controllers.ShowTempo)
	e.POST("produtos/novo", controllers.AddProduto)
	e.DELETE("produtos/deleta/:id", controllers.DelProduto)

	e.GET("/api/v1/produtos", controllers.GetProdutos)
	e.POST("/api/v1/produtos", controllers.StoreProduto)

	e.GET("/tempo-ita", func(c echo.Context) error {
		log.Printf("SSE client connected, ip: %v", c.RealIP())

		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("SSE client disconnected, ip: %v", c.RealIP())
				return nil
			case <-ticker.C:

				// assemble url
				lat, long := -7.381550, -37.185334
				apiUrl := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current=temperature_2m,wind_speed_10m,relative_humidity_2m", lat, long)

				// issue api query
				resp, err := http.Get(apiUrl);
				if err != nil {
					log.Printf("%v", err)
				}
				defer resp.Body.Close()

				// parse payload
				body, _ := io.ReadAll(resp.Body)
				var payload WeatherData
				_ = json.Unmarshal(body, &payload)
				message := fmt.Sprintf("<div>temperatura: %v</div><div>velocidade do vento: %v</div>", payload.Current.Temperature2m, payload.Current.WindSpeed10m)
				

				// fires response
				event := controllers.Event{
					Data: []byte(message),
					Event: []byte("blondel"),
				}
				if err := event.MarshalTo(w); err != nil {
					log.Printf("%v", err)
					return err
				}
				w.Flush()
			}
		}
	})


	e.Logger.Fatal(e.Start(":3853"))
}
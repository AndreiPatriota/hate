package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"hate/handlers"
	"hate/models"
	"hate/types"

	"github.com/labstack/echo/v4"
)




func main() {
	models.InitDb()

	e := echo.New()
	e.Static("/", "assets")

	e.GET("/", handlers.GetIndex)
	e.GET("/sobre/page", handlers.GetSobrePage)
	e.GET("/tempo/page", handlers.GetTempoPage)
	
	e.GET("/produtos/page", handlers.GetProdutosPage)
	e.GET("/produtos/form", handlers.GetProdutosForm)
	e.GET("/produtos", handlers.GetProdutos)
	e.GET("/produtos/:id", handlers.GetProdutosId)
	e.GET("/produtos/form/:id", handlers.GetProdutoFormId)
	e.POST("/produtos", handlers.PostProdutos)
	e.PUT("/produtos/:id", handlers.PutProdutosId)
	e.DELETE("/produtos/:id", handlers.DeletaProdutosId)

	e.GET("/api/v1/produtos", handlers.GetApiProdutos)
	e.POST("/api/v1/produtos", handlers.PostApiProduto)

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
				var payload types.WeatherData
				_ = json.Unmarshal(body, &payload)
				message := fmt.Sprintf("<div>temperatura: %v</div><div>velocidade do vento: %v</div>", payload.Current.Temperature2m, payload.Current.WindSpeed10m)
				

				// fires response
				event := handlers.Event{
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


	e.Logger.Fatal(e.Start(":3000"))
}
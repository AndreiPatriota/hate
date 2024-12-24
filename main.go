package main

import (
	"hate/controllers"
	"hate/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	models.InitDb()

	e := echo.New()
	e.Use(middleware.Static("/static"))

	e.GET("/", controllers.ShowIndex)
	e.GET("/home", controllers.ShowHome)
	e.GET("/produtos", controllers.ShowProdutos)
	e.GET("/sobre", controllers.ShowSobre)
	e.POST("produtos/novo", controllers.AddProduto)
	e.DELETE("produtos/deleta/:id", controllers.DelProduto)

	e.GET("/api/v1/produtos", controllers.GetProdutos)
	e.POST("/api/v1/produtos", controllers.StoreProduto)

	e.Logger.Fatal(e.Start(":3853"))
}
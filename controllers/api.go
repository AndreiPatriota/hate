package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hate/models"
)


func GetProdutos(c echo.Context) error {
	var produtos []models.Produto
	models.DB.Find(&produtos)

	return c.JSON(http.StatusOK, &produtos)
}

func StoreProduto(c echo.Context) error {
	nome := c.FormValue("nome")
	descricao := c.FormValue("descricao")

	p := &models.Produto {
		Nome: nome,
		Descricao: descricao,
	}
	models.DB.Create(p)

	return c.JSON(http.StatusOK, p)
}
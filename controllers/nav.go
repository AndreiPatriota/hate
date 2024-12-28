package controllers

import (
	"hate/models"
	"hate/views/components"
	"hate/views/pages"
	"strconv"

	"github.com/labstack/echo/v4"
)


func ShowIndex(c echo.Context) error {
	return render(c, pages.Index())
}

func ShowHome(c echo.Context) error {
	return render(c, pages.Home())
}
func ShowProdutos(c echo.Context) error {
	var produtos []models.Produto
	models.DB.Find(&produtos)

	return render(c, pages.Produtos(produtos))
}
func ShowSobre(c echo.Context) error {
	return render(c, pages.Sobre())
}
func ShowTempo(c echo.Context) error {
	return render(c, pages.TempoIta())
}

func AddProduto(c echo.Context) error {
	nome := c.FormValue("nome")
	descricao := c.FormValue("descricao")

	p := models.Produto{
		Nome: nome,
		Descricao: descricao,
	}

	models.DB.Create(&p)

	return render(c, components.CardProduto(p))
}
func DelProduto(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	models.DB.Delete(&models.Produto{}, id)

	var produtos []models.Produto
	models.DB.Find(&produtos)

	return render(c, components.ListProdutos(produtos))
}
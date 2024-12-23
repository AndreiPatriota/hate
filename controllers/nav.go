package controllers

import (
	"hate/models"
	"hate/views/pages"

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
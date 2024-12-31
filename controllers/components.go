package controllers

import (
	"hate/models"
	"hate/views/components"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProdutoFormAdd (c echo.Context) error {
	return render(c, components.ProdutoFormAdd())
}

func GetHeader (c echo.Context) error {
	return render(c, components.Header())
}

func GetProdutoList(c echo.Context) error {
	var produtos []models.Produto

	models.DB.Find(&produtos)
	
	return render(c, components.ListProdutos(produtos))
}

func GetCardProduto(c echo.Context) error {
	idRaw := c.Param("id")
	var produto models.Produto
	id, _ := strconv.ParseInt(idRaw, 10, 32) 

	models.DB.Find(&produto, id)

	return render(c, components.CardProduto(produto))
}

func GetProdutoFormEdita(c echo.Context) error {
	idRaw := c.Param("id")
	var produto models.Produto
	id, _ := strconv.ParseInt(idRaw, 10, 32)

	models.DB.Find(&produto, id)

	return render(c, components.ProdutoFormEdita(produto))

}


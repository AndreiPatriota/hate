package handlers

import (
	"hate/models"
	"hate/views/components"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProdutoForm (c echo.Context) error {
	return render(c, components.ProdutoForm())
}

func GetHeader (c echo.Context) error {
	return render(c, components.Header("Titulo"))
}

func GetProdutos(c echo.Context) error {
	var produtos []models.Produto

	models.DB.Find(&produtos)
	
	return render(c, components.Produtos(produtos))
}

func GetProdutosId(c echo.Context) error {
	idRaw := c.Param("id")
	var produto models.Produto
	id, _ := strconv.ParseInt(idRaw, 10, 32) 

	models.DB.Find(&produto, id)

	return render(c, components.ProdutosId(produto))
}

func GetProdutoFormId(c echo.Context) error {
	idRaw := c.Param("id")
	var produto models.Produto
	id, _ := strconv.ParseInt(idRaw, 10, 32)

	models.DB.Find(&produto, id)

	return render(c, components.ProdutosFormId(produto))

}


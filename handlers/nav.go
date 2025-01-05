package handlers

import (
	"hate/models"
	"hate/views/components"
	"hate/views/pages"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)


func GetIndex(c echo.Context) error {
	return render(c, pages.Index("Principal"))
}

func GetHome(c echo.Context) error {
	return render(c, pages.Home())
}
func GetProdutosPage(c echo.Context) error {
	var produtos []models.Produto
	models.DB.Find(&produtos)

	return render(c, pages.ProdutosPage(produtos))
}
func GetSobre(c echo.Context) error {
	return render(c, pages.Sobre())
}
func GetTempo(c echo.Context) error {
	return render(c, pages.TempoIta())
}

func PostProdutos(c echo.Context) error {
	nome := c.FormValue("nome")
	descricao := c.FormValue("descricao")

	p := models.Produto{
		Nome: nome,
		Descricao: descricao,
	}
	models.DB.Create(&p)

	var produtos []models.Produto
	models.DB.Find(&produtos)

	return render(c, components.Produtos(produtos))
}
func PutProdutosId(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	novoNome := c.FormValue("nome")
	novaDescricao := c.FormValue("descricao")

	log.Printf("nome: %v\ndescrição: %v", novoNome, novaDescricao)

	var produto models.Produto
	models.DB.First(&produto, id)

	produto.Nome = novoNome
	produto.Descricao = novaDescricao
	models.DB.Save(&produto)

	return render(c, components.ProdutosId(produto))

}
func DeleteProdutosId(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	models.DB.Delete(&models.Produto{}, id)

	var produtos []models.Produto
	models.DB.Find(&produtos)

	return c.HTML(200, "")
}

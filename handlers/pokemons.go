package handlers

import (
	"fmt"
	"hate/models"
	"hate/views/components"
	"hate/views/pages"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mtslzr/pokeapi-go"
)

func GetPokemonsPage(c echo.Context) error {
	pokemons := []models.Poke{}
	models.DB.Find(&pokemons)


	return render(c, pages.PokemonsPages(pokemons))
} 

func GetPokemonsId(c echo.Context) error {
	var pokemon models.Poke
	idRaw := c.Param("id")
	id, _ := strconv.ParseInt(idRaw, 10, 32) 

	models.DB.Find(&pokemon, id)

	return render(c, components.PokemonsId(pokemon))
}

func PostPokemons(c echo.Context) error {
	nome := c.FormValue("nome")
	var pokemon models.Poke
	p, err := pokeapi.Pokemon(nome)
	log.Printf("BUNDABUNDABUNDABUNDABUNDABUNDAAAAA:   %v", p.ID)
	if err != nil {
		log.Printf("BLONDEL %v", err)
		return c.HTML(404, "")
	}

	pokemon.ID = uint(p.ID)
	pokemon.Nome = p.Name
	pokemon.Tipo = p.Types[0].Type.Name
	pokemon.FotoUrl = p.Sprites.FrontDefault
	pokemon.Peso = float32(p.Weight)
	pokemon.ChoroUrl = fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/cries/main/cries/pokemon/latest/%v.ogg", pokemon.ID)

	// p2, _ := pokeapi.Characteristic(fmt.Sprintf("%v", pokemon.ID))
	// pokemon.Caracteristica = p2.Descriptions[0].Description
	pokemon.Caracteristica = "Blondel"


	models.DB.Create(&pokemon)

	return render(c, components.PokemonsId(pokemon))

}
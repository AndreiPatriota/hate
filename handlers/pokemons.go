package handlers

import (
	"fmt"
	"hate/models"
	"hate/views/components"
	"hate/views/pages"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mtslzr/pokeapi-go"
)

func GetPokemonsPage(c echo.Context) error {
	pokemons := []models.Poke{}
	models.DB.Find(&pokemons)

	return render(c, pages.PokemonsPage(pokemons))
} 


func PostPokemons(c echo.Context) error {
	nome := c.FormValue("nomePokemon")
	nome = strings.ToLower(strings.Trim(nome, " "))
	var pokemon models.Poke
	
	p, err1 := pokeapi.Pokemon(nome)
	p2, err2 := pokeapi.PokemonSpecies(nome)
	if err1 != nil || err2 != nil {
		return c.String(404, fmt.Sprintf("Não achei o Pokemon de nome %s =X", nome))
	}


	pokemon.Caracteristica = p2.FlavorTextEntries[0].FlavorText
	pokemon.Numero = uint64(p.ID)
	pokemon.Nome = p.Name
	pokemon.Tipo = p.Types[0].Type.Name
	pokemon.FotoUrl = p.Sprites.FrontDefault
	pokemon.Peso = float32(p.Weight)
	pokemon.ChoroUrl = fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/cries/main/cries/pokemon/latest/%v.ogg", pokemon.Numero)

	models.DB.Create(&pokemon)

	return render(c, components.PokemonsId(pokemon))
}

func DeletePokemonsId(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	models.DB.Where("Numero = ?", id).Delete(&models.Poke{})

	return c.HTML(200, "")
}
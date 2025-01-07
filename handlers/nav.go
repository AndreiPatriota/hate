package handlers

import (
	"hate/views/pages"

	"github.com/labstack/echo/v4"
)


func GetIndex(c echo.Context) error {
	return render(c, pages.Index("Principal"))
}
func GetHome(c echo.Context) error {
	return render(c, pages.Home())
}
func GetSobre(c echo.Context) error {
	return render(c, pages.Sobre())
}
func GetTempo(c echo.Context) error {
	return render(c, pages.TempoIta())
}


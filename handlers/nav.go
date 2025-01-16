package handlers

import (
	"hate/views/pages"

	"github.com/labstack/echo/v4"
)

func GetIndex(c echo.Context) error {
	return render(c, pages.Home())
}
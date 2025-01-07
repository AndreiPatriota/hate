package handlers

import (
	"hate/views/components"

	"github.com/labstack/echo/v4"
)


func GetHeader (c echo.Context) error {
	return render(c, components.Header("Titulo"))
}




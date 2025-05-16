package routes

import (
	"github/JeffryValle/handlers"

	"github.com/labstack/echo/v4"
)

func ConfigRutas(e *echo.Echo) {
	e.POST("/users", handlers.CrearUsuario)
	e.GET("/users/:id", handlers.GetUserById)
}

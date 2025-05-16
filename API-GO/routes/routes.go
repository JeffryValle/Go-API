package routes

import (
	"github/JeffryValle/handlers"

	"github.com/labstack/echo/v4"
)

func ConfigRutas(e *echo.Echo) {
	e.POST("/users", handlers.CrearUsuario)

	e.DELETE("/users/:id", handlers.DeleteUser)

	e.PUT("/users/:id", handlers.ActualizarUsuario)

	e.GET("/users/:id", handlers.GetUserById)

}

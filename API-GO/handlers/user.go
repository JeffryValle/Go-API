package handlers

import (
	"github/JeffryValle/db"
	"github/JeffryValle/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CrearUsuario(c echo.Context) error {

	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "datos inválidos"})
	}

	query := `insert into users (name, email, password) values (?, ?, ?)`
	result, err := db.DB.Exec(query, req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "No se pudo crear el usuario"})
	}

	id, _ := result.LastInsertId()
	user := models.User{
		ID:    int(id),
		Name:  req.Name,
		Email: req.Email,
	}

	return c.JSON(http.StatusCreated, user)
}

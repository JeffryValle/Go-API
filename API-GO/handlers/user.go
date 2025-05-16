package handlers

import (
	"github/JeffryValle/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CrearUsuario(c echo.Context) error {

	db.Init()
	defer db.CloseConnection()

	req := new(User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "datos inválidos"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al hashear la contraseña"})
	}

	query := `insert into users (name, email, password) values (?, ?, ?)`
	result, err := db.DB.Exec(query, req.Name, req.Email, hash)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "No se pudo crear el usuario"})
	}

	id, _ := result.LastInsertId()
	user := User{
		ID:    int(id),
		Name:  req.Name,
		Email: req.Email,
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUserById(c echo.Context) error {

	db.Init()
	defer db.CloseConnection()

	id := c.Param("id")
	row := db.DB.QueryRow(`select id, name, email from users where id = ?`, id)

	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"Error": "Usuario no encontrado"})
	}

	return c.JSON(http.StatusOK, u)
}

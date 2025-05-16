package handlers

import (
	"fmt"
	"github/JeffryValle/db"
	"net/http"
	"strconv"

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

	fmt.Println("Usuario Creado Con Éxito")
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

	fmt.Println("Usuario Encontrado Con Éxito")
	return c.JSON(http.StatusOK, u)
}

func ActualizarUsuario(c echo.Context) error {

	db.Init()
	defer db.CloseConnection()

	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "datos inválidos"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al hashear la contraseña"})
	}

	query := `update users set name = ?, email = ?, password = ? where id = ?`
	_, err = db.DB.Exec(query, u.Name, u.Email, hash, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo actualizar el usuario"})
	}

	idx, erro := strconv.Atoi(id)
	if erro != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	user := User{
		ID:    idx,
		Name:  u.Name,
		Email: u.Email,
	}

	fmt.Println("Usuario Actualizado Con Éxito")
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	db.Init()
	defer db.CloseConnection()

	id := c.Param("id")
	query := `delete from users where id = ?`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo eliminar el usuario"})
	}

	fmt.Println("Usuario Eliminado Con Éxito")
	return c.JSON(http.StatusOK, echo.Map{"mensaje": "Usuario eliminado Con Éxito"})
}

package main

import (
	"github/JeffryValle/db"
	"github/JeffryValle/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	defer db.CloseConnection()

	e := echo.New()

	routes.ConfigRutas(e)

	e.Logger.Fatal(e.Start(":4000"))
}

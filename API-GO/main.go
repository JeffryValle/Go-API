package main

import (
	"github/JeffryValle/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method:${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))

	routes.ConfigRutas(e)

	// e.Server.Addr = ":4000"

	// e.StartServer(e.Server)

	e.Logger.Fatal(e.Start(":4000"))

}

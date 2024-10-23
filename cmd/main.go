package main

import (
	"net/http"

	"github.com/oik17/mpl-be/internal/database"
	"github.com/oik17/mpl-be/internal/routes"
	"github.com/oik17/mpl-be/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	utils.PrintSiamBanner()
	database.Connect()
	database.RedisConnect()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_custom}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
		CustomTimeFormat: "02/01/2006 15:04:05",
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	routes.TeamRoutes(e)
	routes.UserRoutes(e)
	e.Start(":8080")

}

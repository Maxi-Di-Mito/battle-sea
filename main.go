package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"htmx-app/utils"
)

func main() {
	// db.InitDb()

	server := echo.New()
	server.Use(middleware.Logger())
	server.Static("/static", "static")
	server.Renderer = utils.Temps

	server.GET("/", HomeHandler)

	server.POST("/start-game", InitGameHandler)
	server.POST("/join-game/:gameId", JoinGameHandler)

	server.GET("/game/:id/player/:player", GameHandler)
	server.POST("/click-cell", ClickCellHandler)

	server.GET("/control/:game", SpecHandler)

	server.Logger.Fatal(server.Start(":8765"))
	fmt.Println("RUNNING on 8765")
}

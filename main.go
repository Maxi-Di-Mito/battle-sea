package main

import (
	"fmt"
	"htmx-app/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// db.InitDb()

	server := echo.New()
	server.Use(middleware.Logger())
	server.Static("/static", "static")
	server.Renderer = utils.Temps

	server.GET("/", HomeHandler)

	server.POST("/start-game", InitGameHandler)
	server.GET("/poll-oponent", PollForOponentHandler)
	server.POST("/join-game/:gameId", JoinGameHandler)

	// see your board
	server.GET("/game/:id/player/:player", GameHandler)
	// make a move
	server.POST("/click-cell", ClickCellHandler)

	server.GET("/control/:game", SpecHandler)

	server.Logger.Fatal(server.Start(":8765"))
	fmt.Println("RUNNING on 8765")
}

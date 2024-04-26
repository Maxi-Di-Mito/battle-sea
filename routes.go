package main

import (
	"htmx-app/api/entities"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(ctx echo.Context) error {
	playerNumber := ctx.QueryParam("player")
	var player entities.Player
	if playerNumber == "1" {
		player = *CurrentGame.PlayerOne
	} else {
		player = *CurrentGame.PlayerTwo
	}

	return ctx.Render(http.StatusOK, "index", player)
}

func ClickCell(w http.ResponseWriter, r *http.Request) {

}

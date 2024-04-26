package main

import (
	"htmx-app/api/entities"
	"htmx-app/api/logic"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomeHandler(ctx echo.Context) error {
	playerNumber := ctx.Param("player")
	var player *entities.Player
	if playerNumber == "1" {
		player = logic.CurrentGame.PlayerOne
	} else {
		player = logic.CurrentGame.PlayerTwo
	}

	playerCookie := new(http.Cookie)
	playerCookie.Name = "player"
	playerCookie.Value = playerNumber
	playerCookie.HttpOnly = true
	ctx.SetCookie(playerCookie)

	return ctx.Render(http.StatusOK, "index", player)
}

func ClickCellHandler(ctx echo.Context) error {
	playerHeader := ctx.Request().Header.Get("player")

	attacker, target := logic.GetPlayersFromCookie(playerHeader)
	data := ctx.FormValue("clicked")
	shot := logic.ParseClickedRequest(data)

	modifiedCell := logic.GetShotedCell(attacker, target, shot)

	return ctx.Render(http.StatusOK, "cell", modifiedCell)
}

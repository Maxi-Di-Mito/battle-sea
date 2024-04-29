package main

import (
	"fmt"
	"htmx-app/api/logic"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func HomeHandler(ctx echo.Context) error {
	_, err := ctx.Cookie("playerId")
	if err != nil {
		cookie := new(http.Cookie)
		cookie.Name = "playerId"
		cookie.Value = uuid.New().String()
		cookie.HttpOnly = true
		ctx.SetCookie(cookie)
	}

	return ctx.Render(http.StatusOK, "home", logic.GameList)
}

func InitGameHandler(ctx echo.Context) error {
	playerName := ctx.Request().Form.Get("name")
	playerCookie, err := ctx.Cookie("playerId")
	if err != nil {
		fmt.Println("NO PLAYER ID")
	}
	playerId := playerCookie.Value

	player := logic.GetNewPlayer(playerName, playerId)
	game := logic.InitGame(player)

	ctx.Response().Header().Set("HX-Location", fmt.Sprintf("/game/%s/player/%s", game.ID, player.ID))

	return ctx.String(http.StatusOK, "start")
}

func JoinGameHandler(ctx echo.Context) error {
	playerName := ctx.Request().Form.Get("name")
	gameId := ctx.Param("gameId")
	playerCookie, err := ctx.Cookie("playerId")
	if err != nil {
		fmt.Println("NO PLAYER ID")
	}
	playerId := playerCookie.Value

	player := logic.GetNewPlayer(playerName, playerId)
	game := logic.FindGameById(gameId)

	game.PlayerTwo = player

	game.PlayerOne.Turn = true

	ctx.Response().Header().Set("HX-Location", fmt.Sprintf("/game/%s/player/%s", game.ID, player.ID))

	return ctx.String(http.StatusOK, "join")
}

func GameHandler(ctx echo.Context) error {
	player := logic.GetPlayerFromCookie(ctx)

	return ctx.Render(http.StatusOK, "game", player)
}

func SpecHandler(ctx echo.Context) error {
	gameId := ctx.Param("game")
	game := logic.FindGameById(gameId)
	return ctx.Render(http.StatusOK, "spec", game)
}

func ClickCellHandler(ctx echo.Context) error {
	attacker := logic.GetPlayerFromCookie(ctx)
	target := logic.GetTargetFromCookie(ctx)
	data := ctx.FormValue("clicked")
	shot := logic.ParseClickedRequest(data)

	modifiedCell := logic.GetShotedCell(attacker, target, shot)

	attacker.Turn = false
	target.Turn = true

	return ctx.Render(http.StatusOK, "cell", modifiedCell)
}

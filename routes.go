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

	game.Turn = game.PlayerOne

	ctx.Response().Header().Set("HX-Location", fmt.Sprintf("/game/%s/player/%s", game.ID, player.ID))

	return ctx.String(http.StatusOK, "join")
}

func GameHandler(ctx echo.Context) error {
	boardState := logic.GetStateForPlayerFromCookie(ctx)
	fmt.Println("ESTA ACTIVE", boardState.IsActive())

	return ctx.Render(http.StatusOK, "game", boardState)
}

func PollForOponentHandler(ctx echo.Context) error {
	boardState := logic.GetStateForPlayerFromCookie(ctx)

	if boardState.Oponent == nil {
		return ctx.String(http.StatusOK, "waiting")
	} else {
		ctx.Response().Header().Set("HX-Refresh", "true")
		return ctx.String(http.StatusOK, "oponent joined")
	}
}

func SpecHandler(ctx echo.Context) error {
	gameId := ctx.Param("game")
	game := logic.FindGameById(gameId)
	return ctx.Render(http.StatusOK, "spec", game)
}

func ClickCellHandler(ctx echo.Context) error {
	boardState := logic.GetStateForPlayerFromCookie(ctx)
	attacker := boardState.Player
	target := boardState.Oponent
	data := ctx.FormValue("clicked")
	shot := logic.ParseClickedRequest(data)

	modifiedCell := logic.GetShotedCell(attacker, target, shot)

	boardState.Game.Turn = target

	return ctx.Render(http.StatusOK, "cell", modifiedCell)
}

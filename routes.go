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

// InitGameHandler creates a new game object. Set the game cookie for the player and redirects to the game path
func InitGameHandler(ctx echo.Context) error {
	playerName := ctx.Request().Form.Get("name")
	playerCookie, err := ctx.Cookie("playerId")
	if err != nil {
		fmt.Println("NO PLAYER ID")
	}
	playerId := playerCookie.Value

	player := logic.FindOrCreateNewPlayer(playerName, playerId)
	game := logic.InitGame(player)

	ctx.Response().Header().Set("HX-Location", fmt.Sprintf("/game/%s/player/%s", game.ID, player.ID))

	cookie := new(http.Cookie)
	cookie.Name = "gameId"
	cookie.Value = game.ID
	cookie.HttpOnly = true
	ctx.SetCookie(cookie)
	return ctx.String(http.StatusOK, "start")
}

func JoinGameHandler(ctx echo.Context) error {
	playerName := ctx.Request().Form.Get("name")
	gameId := ctx.Param("gameId")
	playerCookie, err := ctx.Cookie("playerId")

	cookie := new(http.Cookie)
	cookie.Name = "gameId"
	cookie.Value = gameId
	ctx.SetCookie(cookie)

	if err != nil {
		fmt.Println("NO PLAYER ID")
	}
	playerId := playerCookie.Value

	player := logic.FindOrCreateNewPlayer(playerName, playerId)
	game := logic.FindGameById(gameId)

	game.PlayerTwoId = player.ID
	game.PlayerTwoTabs = logic.InitTabs()

	game.Turn = game.PlayerOneId

	ctx.Response().Header().Set("HX-Location", fmt.Sprintf("/game/%s/player/%s", game.ID, player.ID))

	return ctx.String(http.StatusOK, "join")
}

func GameHandler(ctx echo.Context) error {
	gameId := ctx.Param("id")
	playerId := ctx.Param("player")
	cookiePlayerId := logic.GetPlayerIdFromCookie(ctx)
	if playerId != cookiePlayerId {
		return ctx.Render(http.StatusForbidden, "error", "Player does not match game")
	}

	game, err := logic.FindGameAndValidate(gameId, playerId)
	if err != nil {
		ctx.Render(http.StatusNotFound, "error", err.Error())
	}

	player := logic.FindPlayerById(playerId)

	data := logic.GetRenderGameData(game, player)

	fmt.Printf("%v\n", data)
	return ctx.Render(http.StatusOK, "game", data)
}

func PollForOponentHandler(ctx echo.Context) error {
	pId := logic.GetPlayerIdFromCookie(ctx)
	game, err := logic.FindGameAndValidate("id", pId)
	if err != nil {
		return ctx.Render(http.StatusNotFound, "error", "No Game")
	}

	if !game.IsReady() {
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
	gameId := logic.GetGameIdFromCookie(ctx)
	game := logic.FindGameById(gameId)
	pId := logic.GetPlayerIdFromCookie(ctx)
	attacker, target := logic.GetAttackeAndTargetBoards(game, pId)
	data := ctx.FormValue("clicked")
	shot := logic.ParseClickedRequest(data)

	cell := logic.GetShotedCell(attacker, target, shot)

	game.Turn = game.OtherPlayerId(pId)

	return ctx.Render(http.StatusOK, "cell", cell)
}

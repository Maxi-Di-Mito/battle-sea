package logic

import (
	"fmt"
	"htmx-app/api/entities"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func ParseClickedRequest(data string) *entities.ClickedCellRequest {
	coors := strings.Split(data, "-")

	x, err := strconv.Atoi(coors[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coors[1])
	if err != nil {
		panic(err)
	}

	return &entities.ClickedCellRequest{Coor: &entities.Coordinates{X: x, Y: y}}
}

func GetPlayerFromCookie(ctx echo.Context) *entities.Player {
	playerCookie, err := ctx.Cookie("playerId")
	if err != nil {
		fmt.Println("NO PLAYER ID")
	}
	playerId := playerCookie.Value
	fmt.Println("PLAYER ID PARA BUSCAR", playerId)

	for _, game := range GameList {
		if game.PlayerOne != nil && game.PlayerOne.ID == playerId {
			return game.PlayerOne
		} else if game.PlayerTwo != nil && game.PlayerTwo.ID == playerId {
			return game.PlayerTwo
		}
	}
	return nil
}

func GetTargetFromCookie(ctx echo.Context) *entities.Player {
	playerCookie, err := ctx.Cookie("playerId")
	if err != nil {
		fmt.Println("NO PLAYER ID")
	}
	playerId := playerCookie.Value
	fmt.Println("PLAYER ID PARA BUSCAR", playerId)

	for _, game := range GameList {
		if game.PlayerOne != nil && game.PlayerOne.ID == playerId {
			return game.PlayerTwo
		} else if game.PlayerTwo != nil && game.PlayerTwo.ID == playerId {
			return game.PlayerOne
		}
	}
	return nil
}

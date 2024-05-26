package logic

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"htmx-app/api/entities"
	"math/rand"
	"strconv"
	"strings"
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

func GetPlayerIdFromCookie(ctx echo.Context) string {
	playerCookie, err := ctx.Cookie("playerId")
	if err != nil {
		fmt.Println("NO PLAYER ID")
	}
	playerId := playerCookie.Value

	return playerId
}

func FindGameById(id string) *entities.Game {
	for idx, game := range GameList {
		if game.ID == id {
			return &GameList[idx]
		}
	}
	return nil
}

func PopulateBlankBoard(val entities.CellValue) *entities.Board {
	board := entities.Board{Cells: [][]entities.Cell{}}

	for x := 0; x < 10; x++ {
		board.Cells = append(board.Cells, make([]entities.Cell, 10))
		for y := 0; y < 10; y++ {
			currentCell := &board.Cells[x][y]
			currentCell.Coor = &entities.Coordinates{X: x, Y: y}
			currentCell.Value = val
		}
	}

	return &board
}

func PopulateRandomBoats() *entities.Board {
	board := PopulateBlankBoard(entities.CELLVALUE_WATER)

	x := rand.Intn(10)
	y := rand.Intn(10)
	cell := &board.Cells[x][y]
	cell.Value = entities.CELLVALUE_BOAT

	x = rand.Intn(10)
	y = rand.Intn(10)
	cell = &board.Cells[x][y]
	cell.Value = entities.CELLVALUE_BOAT

	return board
}

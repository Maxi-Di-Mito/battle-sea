package logic

import (
	"htmx-app/api/entities"
	"htmx-app/utils"

	"github.com/google/uuid"
)

var currentGame *entities.Game

func GetNewPlayer(name string) *entities.Player {
	player := entities.Player{}

	player.ID = uuid.New().String()
	player.Name = name
	player.AttackTab = PopulateBlankBoard()

	return &player
}

func InitGame(p1 *entities.Player, p2 *entities.Player) *entities.Game {
	game := entities.Game{}

	game.PlayerOne = p1
	game.PlayerTwo = p2

	return &game
}

func PopulateBlankBoard() *entities.Board {
	board := entities.Board{Cells: [][]entities.Cell{}}

	for x := 0; x < 10; x++ {
		board.Cells = append(board.Cells, make([]entities.Cell, 10))
		for y := 0; y < 10; y++ {
			currentCell := &board.Cells[x][y]
			currentCell.Coor = &entities.Coordinates{X: x, Y: y}
			currentCell.Type = utils.CELLTYPE_WATER
			currentCell.Value = "water"
		}
	}

	return &board
}

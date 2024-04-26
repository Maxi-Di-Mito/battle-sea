package logic

import (
	"htmx-app/api/entities"
	"htmx-app/utils"
	"math/rand"

	"github.com/google/uuid"
)

var CurrentGame *entities.Game

func GetNewPlayer(name string) *entities.Player {
	player := entities.Player{}

	player.ID = uuid.New().String()
	player.Name = name
	player.AttackTab = PopulateBlankBoard(utils.CELLVALUE_UNKNOWN)
	player.HomeTab = PopulateRandomBoats()

	return &player
}

func InitGame(p1 *entities.Player, p2 *entities.Player) *entities.Game {
	game := entities.Game{}

	game.PlayerOne = p1
	game.PlayerTwo = p2

	return &game
}

func PopulateBlankBoard(val utils.CellValue) *entities.Board {
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
	board := PopulateBlankBoard(utils.CELLVALUE_WATER)

	x := rand.Intn(10)
	y := rand.Intn(10)
	cell := &board.Cells[x][y]
	cell.Value = utils.CELLVALUE_BOAT

	x = rand.Intn(10)
	y = rand.Intn(10)
	cell = &board.Cells[x][y]
	cell.Value = utils.CELLVALUE_BOAT

	return board
}

func GetShotedCell(attacker *entities.Player, target *entities.Player, shot *entities.ClickedCellRequest) *entities.Cell {
	shotted := &target.HomeTab.Cells[shot.Coor.X][shot.Coor.Y]
	marker := &attacker.AttackTab.Cells[shot.Coor.X][shot.Coor.Y]

	if shotted.Value == utils.CELLVALUE_BOAT {
		shotted.Value = utils.CELLVALUE_DEAD
		marker.Value = utils.CELLVALUE_DEAD
	} else if shotted.Value == utils.CELLVALUE_WATER {
		marker.Value = utils.CELLVALUE_WATER
	}

	return marker
}

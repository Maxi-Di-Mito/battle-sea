package logic

import (
	"htmx-app/api/entities"
	"math/rand"

	"github.com/google/uuid"
)

var GameList []entities.Game

func GetNewPlayer(name string, id string) *entities.Player {
	player := entities.Player{}

	player.ID = id
	player.Name = name
	player.AttackTab = PopulateBlankBoard(entities.CELLVALUE_UNKNOWN)
	player.HomeTab = PopulateRandomBoats()

	return &player
}

func InitGame(p1 *entities.Player) *entities.Game {
	game := entities.Game{}

	game.ID = uuid.New().String()
	game.PlayerOne = p1

	GameList = append(GameList, game)

	return &game
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

func GetShotedCell(attacker *entities.Player, target *entities.Player, shot *entities.ClickedCellRequest) *entities.Cell {
	shotted := &target.HomeTab.Cells[shot.Coor.X][shot.Coor.Y]
	marker := &attacker.AttackTab.Cells[shot.Coor.X][shot.Coor.Y]

	if shotted.Value == entities.CELLVALUE_BOAT {
		shotted.Value = entities.CELLVALUE_DEAD
		marker.Value = entities.CELLVALUE_DEAD
	} else if shotted.Value == entities.CELLVALUE_WATER {
		marker.Value = entities.CELLVALUE_WATER
	}

	return marker
}

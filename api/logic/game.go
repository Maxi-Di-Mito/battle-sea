package logic

import (
	"errors"
	"htmx-app/api/entities"

	"github.com/google/uuid"
)

var GameList []entities.Game

func InitGame(p1 *entities.Player) *entities.Game {
	game := entities.Game{}

	game.ID = uuid.New().String()
	game.PlayerOneId = p1.ID
	game.PlayerOneTabs = InitTabs()

	GameList = append(GameList, game)

	return &game
}

func InitTabs() *entities.GameTabs {
	tabs := entities.GameTabs{}

	tabs.AttackTab = PopulateBlankBoard(entities.CELLVALUE_WATER)
	tabs.HomeTab = PopulateRandomBoats()

	return &tabs
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

// Finds a game and validates that the player is on that game
func FindGameAndValidate(gameId string, playerId string) (*entities.Game, error) {
	game := FindGameById(gameId)
	if game == nil {
		return nil, errors.New("Game not found")
	}

	if game.PlayerOneId != playerId && game.PlayerTwoId != playerId {
		return nil, errors.New("Game and Player do not match")
	}

	return game, nil
}

func GetRenderGameData(game *entities.Game, player *entities.Player) *entities.GameRenderData {
	data := entities.GameRenderData{}
	data.Game = game
	if game.PlayerOneId == player.ID {
		data.Tabs = game.PlayerOneTabs
	} else {
		data.Tabs = game.PlayerTwoTabs
	}

	data.IsActive = game.Turn == player.ID

	return &data
}

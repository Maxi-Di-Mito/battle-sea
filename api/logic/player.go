package logic

import "htmx-app/api/entities"

var PlayerList []entities.Player

func FindOrCreateNewPlayer(name string, id string) *entities.Player {
	player := FindPlayerById(id)
	if player != nil {
		return player
	}
	newPlayer := entities.Player{}

	newPlayer.ID = id
	newPlayer.Name = name

	PlayerList = append(PlayerList, newPlayer)

	return &newPlayer
}

func FindPlayerById(id string) *entities.Player {
	for index, player := range PlayerList {
		if player.ID == id {
			return &PlayerList[index]
		}
	}

	return nil
}

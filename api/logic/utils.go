package logic

import (
	"fmt"
	"htmx-app/api/entities"
	"strconv"
	"strings"
)

func ParseClickedRequest(data string) *entities.ClickedCellRequest {
	coors := strings.Split(data, "-")
	fmt.Println("------------- where shot", coors)

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

func GetPlayersFromCookie(number string) (*entities.Player, *entities.Player) {
	fmt.Println("------------- who shot", number)
	if number == "1" {
		return CurrentGame.PlayerOne, CurrentGame.PlayerTwo
	} else {
		return CurrentGame.PlayerTwo, CurrentGame.PlayerOne
	}
}

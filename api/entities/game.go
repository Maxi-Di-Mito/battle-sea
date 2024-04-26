package entities

import "htmx-app/utils"

type Game struct {
	PlayerOne *Player
	PlayerTwo *Player
}

type Player struct {
	ID        string
	Name      string
	AttackTab *Board
	HomeTab   *Board
	Boats     []Boat
}

type Boat struct {
	Name      string
	Length    int
	Position  *Coordinates
	Direction rune
}

type Coordinates struct {
	X int
	Y int
}

type Board struct {
	Cells [][]Cell
}

type Cell struct {
	Coor  *Coordinates
	Value utils.CellValue
}

type ClickedCellRequest struct {
	Coor  *Coordinates
	Value utils.CellValue
}

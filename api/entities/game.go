package entities

type Game struct {
	PlayerOne *Player
	PlayerTwo *Player
	Turn      *Player
	ID        string
}

type Player struct {
	ID        string
	Name      string
	AttackTab *Board
	HomeTab   *Board
	Boats     []Boat
	Turn      bool
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
	Value CellValue
}

type ClickedCellRequest struct {
	Coor  *Coordinates
	Value CellValue
}

type BoardState struct {
	Player  *Player
	Oponent *Player
	Game    *Game
}

func (board *BoardState) IsActive() bool {
	if board.Game.Turn == nil {
		return false
	} else {
		return board.Player.ID == board.Game.Turn.ID
	}
}

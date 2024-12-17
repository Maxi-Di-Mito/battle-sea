package entities

type Game struct {
	PlayerOneId   string
	PlayerOneTabs *GameTabs
	PlayerTwoId   string
	PlayerTwoTabs *GameTabs
	Turn          string
	ID            string
}

func (game *Game) IsReady() bool {
	return game.PlayerOneId != "" && game.PlayerTwoId != ""
}

func (game *Game) OtherPlayerId(pId string) string {
	if game.PlayerOneId == pId {
		return game.PlayerTwoId
	} else {
		return game.PlayerOneId
	}
}

type GameTabs struct {
	AttackTab *Board
	HomeTab   *Board
	Boats     []Boat
}

type Player struct {
	ID   string
	Name string
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

// The data needed to render a game for a player
type GameRenderData struct {
	Game     *Game
	Tabs     *GameTabs
	IsActive bool
	Ready    bool
}

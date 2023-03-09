package game

type CellState int

const (
	Open CellState = iota
	Closed
)

type CellType int

const (
	Normal CellType = iota
	BHole
)

type GameState int

const (
	Uninit GameState = iota
	Ongoing
	Won
	Lost
)

type Cell struct {
	state    CellState
	adjBHCnt uint
}
package game

type CellState int

const (
	Open CellState = iota
	Closed
)

type CellType int

const (
	Normal CellType = iota
	BlackHole
)

type GameState int

const (
	Uninit GameState = iota
	Ongoing
	Won
	Lost
)

package game

type LayoutProvider interface {
	GetBoardLayout(w, h, hnum uint) ([][]uint8, error)
}

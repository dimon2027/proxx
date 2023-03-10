package game

type LayoutProvider interface {
	GetBoardLayout(w, h, BHCnt int) ([][]uint8, error)
}

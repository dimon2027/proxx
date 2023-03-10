package game

type Cell struct {
	x, y     int
	s        CellState
	t        CellType
	adjBHCnt int
}

func NewCell(x, y int) Cell {
	return Cell{
		x: x,
		y: y,
		s: Closed,
		t: Normal,
	}
}

func (c *Cell) GetType() CellType {
	return c.t
}

func (c *Cell) GetState() CellState {
	return c.s
}

func (c *Cell) GetAdjBHolesCnt() int {
	return c.adjBHCnt
}

func (c *Cell) setType(t CellType) {
	c.t = t
}

func (c *Cell) setState(s CellState) {
	c.s = s
}

func (c *Cell) setAdjBHolesCnt(cnt int) {
	c.adjBHCnt = cnt
}

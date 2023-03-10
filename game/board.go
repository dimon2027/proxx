package game

type board struct {
	w, h  int
	cells [][]Cell
}

func newBoard() *board {
	return &board{}
}

func (b *board) Init(w, h int) {
	b.w = w
	b.h = h

	cells := make([][]Cell, w)
	for i := range cells {
		cells[i] = make([]Cell, h)
		for j := range cells[i] {
			cells[i][j] = NewCell(i, j)
		}
	}

	b.cells = cells
}

// TODO: consider using adjacency matrix here instead
func (b *board) GetAdjCells(c *Cell) []*Cell {
	if b == nil {
		return nil
	}

	if c == nil {
		return nil
	}

	// Optimize this
	a := make([]*Cell, 0, 8)

	for i := c.x - 1; i <= c.x+1; i++ {
		for j := c.y - 1; j <= c.y+1; j++ {
			if i >= 0 && i < b.w &&
				j >= 0 && j < b.h &&
				!(i == c.x && j == c.y) {
				a = append(a, b.GetCell(i, j))
			}
		}
	}

	return a
}

func (b *board) GetCell(x, y int) *Cell {
	if x >= b.w || y >= b.h {
		return nil
	}

	return &(b.cells[x][y])
}

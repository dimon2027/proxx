package game

import (
	"errors"
)

const (
	MaxBoardW          = 40
	MaxBoardH          = 40
	MinBoardW          = 1
	MinBoardH          = 1
	MinCellCount       = 2
	MinBlackHolesCount = 1
)

type Game struct {
	state   GameState
	board   *board
	w, h    int
	BHCnt   int
	openCnt int
}

func New() *Game {
	return &Game{
		state: Uninit,
	}
}

func (g *Game) Init(w, h, BHCnt int, lp LayoutProvider) error {
	if w < MinBoardW || h < MinBoardH || w > MaxBoardW || h > MaxBoardH || w*h < MinCellCount {
		return errors.New("board dimenisons are invalid")
	}

	if BHCnt < MinBlackHolesCount || BHCnt >= w*h {
		return errors.New("black holes count is invalid")
	}

	g.w = w
	g.h = h
	g.BHCnt = BHCnt

	g.board = newBoard()
	g.board.Init(w, h)

	g.putBlackHoles(lp)
	g.putNumbers()

	g.state = Ongoing

	return nil
}

func (g *Game) MakeTurn(x, y int) (GameState, error) {
	if g.state != Ongoing {
		return g.state, errors.New("invalid state")
	}

	if x >= g.w || y >= g.h || x < 0 || y < 0 {
		return g.state, errors.New("invalid input")
	}

	c := g.board.GetCell(x, y)

	if c.GetType() == BlackHole {
		g.state = Lost
		return g.state, nil
	}

	g.openCell(c)

	// If the number of closed cells is equal to the number of bombs - it's a victory
	if g.w*g.h-g.openCnt == g.BHCnt {
		g.state = Won
		return g.state, nil
	}

	return g.state, nil
}

// GetBoard retruns a *read-only* cell matrix for consumption by UI component
func (g *Game) GetBoard() [][]Cell {
	return g.board.cells
}

func (g *Game) GetState() GameState {
	return g.state
}

func (g *Game) putBlackHoles(lp LayoutProvider) error {
	layout, err := lp.GetBoardLayout(g.w, g.h, g.BHCnt)
	if err != nil {
		return err
	}

	for i := range layout {
		for j := range layout[0] {
			if layout[i][j] == 1 {
				g.board.GetCell(i, j).setType(BlackHole)
			}
		}
	}

	return nil
}

func (g *Game) putNumbers() {
	for i := 0; i < g.w; i++ {
		for j := 0; j < g.h; j++ {
			cell := g.board.GetCell(i, j)
			if cell.GetType() == BlackHole {
				for _, c := range g.board.GetAdjCells(cell) {
					if c.GetType() == Normal {
						c.setAdjBHolesCnt(c.GetAdjBHolesCnt() + 1)
					}
				}
			}
		}
	}
}

func (g *Game) openCell(cell *Cell) {
	if cell.GetState() == Open ||
		cell.GetType() != Normal {
		return
	}

	cell.setState(Open)
	g.openCnt++

	if cell.adjBHCnt != 0 {
		return
	}

	var stack []*Cell
	stack = append(stack, cell)

	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, ac := range g.board.GetAdjCells(c) {
			if ac.GetState() == Closed {
				ac.setState(Open)
				g.openCnt++
				if ac.adjBHCnt == 0 {
					stack = append(stack, ac)
				}
			}
		}
	}
}

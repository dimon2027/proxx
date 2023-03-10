package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestLayoutProvider struct {
	layout [][]uint8
}

func NewTestLayoutProvider(layout [][]uint8) *TestLayoutProvider {
	return &TestLayoutProvider{
		layout: layout,
	}
}
func (lp *TestLayoutProvider) GetBoardLayout(w, h, BHCnt int) ([][]uint8, error) {
	return lp.layout, nil
}

func TestGameInit(t *testing.T) {
	w := 3
	h := 3
	bhnum := 2
	layout := [][]uint8{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 0}}

	expected := [][]Cell{
		{{0, 0, Closed, Normal, 0}, {0, 1, Closed, Normal, 2}, {0, 2, Closed, BlackHole, 0}},
		{{1, 0, Closed, Normal, 0}, {1, 1, Closed, Normal, 2}, {1, 2, Closed, BlackHole, 0}},
		{{2, 0, Closed, Normal, 0}, {2, 1, Closed, Normal, 1}, {2, 2, Closed, Normal, 1}}}

	g := New()
	g.Init(w, h, bhnum, NewTestLayoutProvider(layout))

	assert.Equal(t, expected, g.board.cells)
}

func TestOpenCellNotAdj(t *testing.T) {
	w := 3
	h := 3
	bhnum := 2
	layout := [][]uint8{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 0}}

	x := 1
	y := 0

	expected := [][]Cell{
		{{0, 0, Open, Normal, 0}, {0, 1, Open, Normal, 2}, {0, 2, Closed, BlackHole, 0}},
		{{1, 0, Open, Normal, 0}, {1, 1, Open, Normal, 2}, {1, 2, Closed, BlackHole, 0}},
		{{2, 0, Open, Normal, 0}, {2, 1, Open, Normal, 1}, {2, 2, Closed, Normal, 1}}}

	g := New()
	g.Init(w, h, bhnum, NewTestLayoutProvider(layout))
	g.openCell(g.board.GetCell(x, y))

	assert.Equal(t, expected, g.board.cells)
}

func TestOpenCellAdj(t *testing.T) {
	w := 3
	h := 3
	bhnum := 2
	layout := [][]uint8{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 0}}

	x := 0
	y := 1

	expected := [][]Cell{
		{{0, 0, Closed, Normal, 0}, {0, 1, Open, Normal, 2}, {0, 2, Closed, BlackHole, 0}},
		{{1, 0, Closed, Normal, 0}, {1, 1, Closed, Normal, 2}, {1, 2, Closed, BlackHole, 0}},
		{{2, 0, Closed, Normal, 0}, {2, 1, Closed, Normal, 1}, {2, 2, Closed, Normal, 1}}}

	g := New()
	g.Init(w, h, bhnum, NewTestLayoutProvider(layout))
	g.openCell(g.board.GetCell(x, y))

	assert.Equal(t, expected, g.board.cells)
}

func TestInvalidBoardDimensions(t *testing.T) {
	g1 := New()
	err := g1.Init(1, 1, 1, NewRandomLayout())
	assert.Error(t, err)

	g2 := New()
	err = g2.Init(-2, 2, 1, NewRandomLayout())
	assert.Error(t, err)
}

func TestInvalidBlackHolesCnt(t *testing.T) {
	g1 := New()
	err := g1.Init(5, 5, 0, NewRandomLayout())
	assert.Error(t, err)

	g2 := New()
	err = g2.Init(5, 5, 25, NewRandomLayout())
	assert.Error(t, err)
}

func TestMakeTurnLosing(t *testing.T) {
	w := 3
	h := 3
	bhnum := 2
	layout := [][]uint8{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 0}}

	x := 0
	y := 2

	g := New()
	g.Init(w, h, bhnum, NewTestLayoutProvider(layout))

	s, err := g.MakeTurn(x, y)
	assert.NoError(t, err)
	assert.Equal(t, Lost, s)
}

func TestMakeTurnWinning(t *testing.T) {
	w := 3
	h := 3
	bhnum := 2
	layout := [][]uint8{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 0}}

	g := New()
	g.Init(w, h, bhnum, NewTestLayoutProvider(layout))

	s, err := g.MakeTurn(2, 0)
	assert.NoError(t, err)
	assert.Equal(t, Ongoing, s)

	s, err = g.MakeTurn(2, 1)
	assert.NoError(t, err)
	assert.Equal(t, Ongoing, s)

	s, err = g.MakeTurn(2, 2)
	assert.NoError(t, err)
	assert.Equal(t, Won, s)
}

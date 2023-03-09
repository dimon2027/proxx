package game

import "github.com/dimon2027/proxx/game/layouts/random"

type Game struct {
	w       uint
	h       uint
	hnum    uint
	opennum uint
	state   GameState
	board   [][]Cell
}

func New() *Game {
	return &Game{
		state: Uninit,
	}
}

func (g *Game) Init(w, h, hnum uint) {
	// TODO: validate input
	g.w = w
	g.h = h
	g.hnum = hnum
	// Place bombs
	g.placeBombs()
}

func (g *Game) placeBombs() {
	random.New()
}

package game

import (
	"errors"
	"math/rand"
	"time"
)

const (
	MinW = 1
	MinH = 1
	MaxW = 100000
	MaxH = 100000
)

type RandomLayout struct {
}

func NewRandomLayout() *RandomLayout {
	return &RandomLayout{}
}

func (l *RandomLayout) GetBoardLayout(w, h, BHCnt int) ([][]uint8, error) {
	if w < MinW || h < MinH || w > MaxW || h > MaxH {
		return nil, errors.New("invalid dimensions")
	}

	if BHCnt < 0 || BHCnt > w*h {
		return nil, errors.New("number of black holes number is invalid")
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	perms := r.Perm(w * h)

	layout := make([][]uint8, h)

	for i := 0; i < h; i++ {
		layout[i] = make([]uint8, w)
		for j := 0; j < int(w); j++ {
			if perms[i*w+j] < BHCnt {
				layout[i][j] = 1
			} else {
				layout[i][j] = 0
			}
		}
	}

	return layout, nil
}

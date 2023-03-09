package randomlayout

import (
	"errors"
	"math/rand"
	"time"
)

type Layout struct {
}

func New() *Layout {
	return &Layout{}
}

// GetBoardLayout returns w by h matrix where each cell is either 1 or 0.
// 1 - corresponding cell contains a bomb,
// 0 - corresponding cell does not contain a bomb.
func (l *Layout) GetBoardLayout(w, h, hnum uint) ([][]uint8, error) {
	// TODO: Validate input
	if hnum > w*h {
		return nil, errors.New("number of black holes can't be greater than the total number of cells")
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	perms := r.Perm(int(w * h))

	layout := make([][]uint8, h)

	for i := 0; i < int(h); i++ {
		layout[i] = make([]uint8, w)
		for j := 0; j < int(w); j++ {
			if perms[i*int(w)+j] < int(hnum) {
				layout[i][j] = 1
			} else {
				layout[i][j] = 0
			}
		}
	}

	return layout, nil
}

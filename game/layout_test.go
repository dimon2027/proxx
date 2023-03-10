package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomLayout(t *testing.T) {
	w := 100
	h := 50
	hnum := 7

	p := NewRandomLayout()

	layout, err := p.GetBoardLayout(w, h, hnum)
	assert.NoError(t, err)

	assert.Equal(t, len(layout), h)
	assert.Equal(t, len(layout[0]), w)

	var hnumActual int
	for i := range layout {
		for j := range layout[i] {
			if layout[i][j] == 1 {
				hnumActual++
			}
		}
	}

	assert.Equal(t, hnum, hnumActual)
}

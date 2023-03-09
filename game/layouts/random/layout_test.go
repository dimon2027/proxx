package randomlayout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	p := New()
	w := 100
	h := 50
	hnum := 7

	layout, err := p.GetBoardLayout(uint(w), uint(h), uint(hnum))
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

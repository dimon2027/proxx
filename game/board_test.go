package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdjCells(t *testing.T) {
	w := 5
	h := 5

	b := newBoard()
	b.Init(w, h)

	// Upper left corner
	ac := b.GetAdjCells(b.GetCell(0, 0))
	assert.Equal(t, 3, len(ac))

	// Upper right corner
	ac = b.GetAdjCells(b.GetCell(w-1, 0))
	assert.Equal(t, 3, len(ac))

	// Bottom left corner
	ac = b.GetAdjCells(b.GetCell(0, h-1))
	assert.Equal(t, 3, len(ac))

	// Upper right corner
	ac = b.GetAdjCells(b.GetCell(w-1, h-1))
	assert.Equal(t, 3, len(ac))

	// Upper mid
	ac = b.GetAdjCells(b.GetCell(w/2, 0))
	assert.Equal(t, 5, len(ac))

	// Right mid
	ac = b.GetAdjCells(b.GetCell(w-1, h/2))
	assert.Equal(t, 5, len(ac))

	// Bottom mid
	ac = b.GetAdjCells(b.GetCell(w/2, h-1))
	assert.Equal(t, 5, len(ac))

	// Left mid
	ac = b.GetAdjCells(b.GetCell(0, h/2))
	assert.Equal(t, 5, len(ac))

	// Center
	ac = b.GetAdjCells(b.GetCell(w/2, h/2))
	assert.Equal(t, 8, len(ac))
}

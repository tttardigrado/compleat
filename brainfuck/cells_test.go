package brainfuck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCellsFromIntSlice(t *testing.T) {
	got := cellsFromIntSlice([]int{1, 2, -3, 55})

	want := CellList{Cell(1), Cell(2), Cell(-3), Cell(55)}

	assert.Equal(t, got, want)
}

func TestIncrementCell(t *testing.T) {
	t.Run("increment a cell", func(t *testing.T) {
		got := Cell(10)

		got.Increment()

		want := Cell(11)

		assert.Equal(t, got, want)
	})

	t.Run("increment a cell on the edge (255) ", func(t *testing.T) {
		got := Cell(255)

		got.Increment()

		want := Cell(0)

		assert.Equal(t, got, want)
	})
}

func TestDecrementCell(t *testing.T) {
	t.Run("increment a cell", func(t *testing.T) {
		got := Cell(10)

		got.Decrement()

		want := Cell(9)

		assert.Equal(t, got, want)
	})

	t.Run("increment a cell on the edge (255) ", func(t *testing.T) {
		got := Cell(0)

		got.Decrement()

		want := Cell(255)

		assert.Equal(t, got, want)
	})
}

func TestIsNull(t *testing.T) {
	t.Run("null cell", func(t *testing.T) {
		c := Cell(0)

		got := c.IsNull()

		want := true

		assert.Equal(t, got, want)
	})

	t.Run("not null", func(t *testing.T) {
		c := Cell(5)

		got := c.IsNull()

		want := false

		assert.Equal(t, got, want)
	})
}

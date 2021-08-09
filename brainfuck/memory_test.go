package brainfuck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCell(t *testing.T) {

	t.Run("non-empty memory", func(t *testing.T) {
		memory := NewMemory([]int{10, 20, 30}, 2)

		memory.NewCell()

		want := NewMemory([]int{10, 20, 30, 0}, 2)

		assert.Equal(t, memory, want)
	})

	t.Run("empty memory", func(t *testing.T) {
		memory := NewMemory([]int{}, 0)

		memory.NewCell()

		want := NewMemory([]int{0}, 0)

		assert.Equal(t, memory, want)
	})

}

func TestMove(t *testing.T) {

	t.Run("move +1", func(t *testing.T) {

		memory := NewMemory([]int{10, 20}, 0)

		memory.MoveUp()

		want := NewMemory([]int{10, 20}, 1)

		assert.Equal(t, memory, want)

	})

	t.Run("move +1 at the edge", func(t *testing.T) {
		memory := NewMemory([]int{10, 20}, 1)

		memory.MoveUp()

		want := NewMemory([]int{10, 20, 0}, 2)

		assert.Equal(t, memory, want)
	})

	t.Run("move -1", func(t *testing.T) {
		memory := NewMemory([]int{10, 20}, 1)

		err := memory.MoveDown()

		if err != nil {
			t.Errorf("Got an error moving the cursor down")
		}

		want := NewMemory([]int{10, 20}, 0)

		assert.Equal(t, memory, want)
	})

	t.Run("move -1 at the edge", func(t *testing.T) {
		memory := NewMemory([]int{10, 20}, 0)

		err := memory.MoveDown()

		if err == nil {
			t.Errorf("Did not get an error, but should")
		}

	})

}

func TestCurrentCell(t *testing.T) {
	t.Run("get a cell", func(t *testing.T) {
		memory := NewMemory([]int{10, 20, 30, 40, 50}, 1)

		got := memory.CurrentCell()
		want := &memory.cells[1]

		assert.Equal(t, got, want)
	})

	t.Run("get cell cursor > len", func(t *testing.T) {
		memory := NewMemory([]int{10, 20}, 2)

		got := memory.CurrentCell()
		want := &memory.cells[1]

		assert.Equal(t, got, want)

	})

}

func TestIncrement(t *testing.T) {

	t.Run("increment 1", func(t *testing.T) {
		memory := NewMemory([]int{10, 20, 30}, 1)

		memory.Increment()

		want := NewMemory([]int{10, 21, 30}, 1)

		assert.Equal(t, memory, want)
	})

	t.Run("increment 1 at 255", func(t *testing.T) {
		memory := NewMemory([]int{10, 255, 30}, 1)

		memory.Increment()

		want := NewMemory([]int{10, 0, 30}, 1)

		assert.Equal(t, memory, want)
	})

}

func TestDecrement(t *testing.T) {

	t.Run("decrement 1", func(t *testing.T) {
		memory := NewMemory([]int{10, 20, 30}, 1)

		memory.Decrement()

		want := NewMemory([]int{10, 19, 30}, 1)

		assert.Equal(t, memory, want)
	})

	t.Run("decrement 1 at 0", func(t *testing.T) {
		memory := NewMemory([]int{10, 0, 30}, 1)

		memory.Decrement()

		want := NewMemory([]int{10, 255, 30}, 1)

		assert.Equal(t, memory, want)
	})

}

func TestIsNullMemory(t *testing.T) {

	t.Run("null", func(t *testing.T) {
		memory := NewMemory([]int{10, 0, 30}, 1)

		got := memory.IsCurrentNull()

		want := true

		assert.Equal(t, got, want)
	})

	t.Run("not null", func(t *testing.T) {
		memory := NewMemory([]int{10, 20, 30}, 1)

		got := memory.IsCurrentNull()

		want := false

		assert.Equal(t, got, want)
	})

}

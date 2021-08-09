package brainfuck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgram(t *testing.T) {
	got := NewProgram([]rune{'>', '+', '-', '<'})

	want := Program{
		Commands: []rune{'>', '+', '-', '<'},
		Cursor:   0,
		Mem:      NewMemory([]int{0}, 0)}

	assert.Equal(t, got, want)
}

func TestIncrementCursor(t *testing.T) {
	got := NewProgram([]rune{'>', '+', '+'})

	got.IncrementCursor()

	want := Program{
		Commands: []rune{'>', '+', '+'},
		Cursor:   1,
		Mem:      NewMemory([]int{0}, 0)}

	assert.Equal(t, got, want)
}

func TestDecrementCursor(t *testing.T) {
	got := Program{
		Commands: []rune{'>', '+', '+'},
		Cursor:   2,
		Mem:      NewMemory([]int{0}, 0)}

	got.DecrementCursor()

	want := Program{
		Commands: []rune{'>', '+', '+'},
		Cursor:   1,
		Mem:      NewMemory([]int{0}, 0)}

	assert.Equal(t, got, want)
}

func TestLength(t *testing.T) {
	t.Run("empty program", func(t *testing.T) {
		p := NewProgram([]rune{})

		got := p.Length()

		want := 0

		assert.Equal(t, got, want)
	})
	t.Run("non-empty program", func(t *testing.T) {
		p := NewProgram([]rune{'>', '+', '-', '2'})

		got := p.Length()

		want := 4

		assert.Equal(t, got, want)
	})

}

func TestGetCurrentCommand(t *testing.T) {
	t.Run("normal cursor", func(t *testing.T) {
		p := Program{
			Commands: []rune{'>', '-', '+'},
			Cursor:   1,
			Mem:      NewMemory([]int{0}, 0)}

		got, err := p.GetCurrentCommand()

		if err != nil {
			t.Errorf("Got an error that was not expected. Info: %v", got)
		}

		want := '-'

		assert.Equal(t, got, want)
	})

	t.Run("edge cursor", func(t *testing.T) {
		p := Program{
			Commands: []rune{'>', '-', '+'},
			Cursor:   2,
			Mem:      NewMemory([]int{0}, 0)}

		got, err := p.GetCurrentCommand()

		if err != nil {
			t.Errorf("Got an error that was not expected. Info: %v", got)
		}

		want := '+'

		assert.Equal(t, got, want)
	})

	t.Run("larger than length cursor", func(t *testing.T) {
		p := Program{
			Commands: []rune{'>', '-', '+'},
			Cursor:   5,
			Mem:      NewMemory([]int{0}, 0)}

		_, err := p.GetCurrentCommand()

		if err == nil {
			t.Error("Did not get an error but expected one")
		}
	})

	t.Run("negative cursor", func(t *testing.T) {
		p := Program{
			Commands: []rune{'>', '-', '+'},
			Cursor:   5,
			Mem:      NewMemory([]int{0}, 0)}

		_, err := p.GetCurrentCommand()

		if err == nil {
			t.Error("Did not get an error but expected one")
		}
	})

}

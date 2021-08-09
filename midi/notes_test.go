package midipackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNote(t *testing.T) {
	t.Run("test negative length", func(t *testing.T) {
		got := NewNote(1, 100, -55)
		want := Note{Index: 1,
			Octave:     3,
			Velocity:   100,
			LenDivider: 55}
		assert.Equal(t, want, got)
	})

	t.Run("test negative note", func(t *testing.T) {

		// D2
		got := NewNote(-10, 100, 2)

		want := Note{Index: 2,
			Octave:     2,
			Velocity:   100,
			LenDivider: 2}
		assert.Equal(t, want, got)
	})

}

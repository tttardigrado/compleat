package converter

import (
	"testing"

    m "compleat/midipackage"
	"github.com/stretchr/testify/assert"
)

func TestProcessBF(t *testing.T){
	t.Run("C major", func(t *testing.T) {
		got := ProcessBF([]rune{'+','-', '#', 'd', '.', '\n', '[', ']', '.', ',', '>', '<'}, m.CM)
		want := m.MusicProgram{m.D3, m.C3, m.G3, m.A3, m.B3, m.C4, m.F3, m.E3}

		assert.Equal(t, want, got)
	})
	t.Run("A minor", func(t *testing.T) {
		got := ProcessBF([]rune{'-', '+', '<', '>', '[', ']', '.', ','}, m.Am)

		want := m.MusicProgram{m.A3, m.B3, m.C4, m.D4, m.E4, m.F4, m.G4, m.A4}

		assert.Equal(t, want, got)
	})


}

func TestProcessMidi(t *testing.T){
	t.Run("C major", func(t *testing.T) {
		got := ProcessMidi(m.MusicProgram{m.D3, m.C3, m.G3, m.A3, m.B3, m.C4, m.F3, m.E3}, m.CM)

		want := []rune{'+','-' , '[', ']', '.', ',', '>', '<'}

		assert.Equal(t, want, got)
	})

	t.Run("A minor", func(t *testing.T) {
		got := ProcessMidi( m.MusicProgram{m.A3, m.B3, m.C4, m.D4, m.E4, m.F4, m.G4, m.A4}, m.Am)

		want :=[]rune{'-', '+', '<', '>', '[', ']', '.', ','}

		assert.Equal(t, want, got)
	})


}

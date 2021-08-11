package midipackage

import (
	r "gitlab.com/gomidi/midi/reader"
)

func (m *MusicProgram) noteOn(p *r.Position, channel, key, vel uint8) {
	*m = append(*m, MidiNote(key))
}

func Read(file string) (MusicProgram, error) {
	var m MusicProgram
	rd := r.New(r.NoLogger(), r.NoteOn(m.noteOn))
	err := r.ReadSMFFile(rd, file)
	return m, err
}

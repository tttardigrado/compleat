package midipackage

import (
	r "gitlab.com/gomidi/midi/reader"
)

// Function that is called when the Reader detects the beggining of a new note.
func (m *MusicProgram) noteOn(p *r.Position, channel, key, vel uint8) {
	*m = append(*m, MidiNote(key))
}

// Read a file and return it as a Music Program (colection of midi notes)
func Read(file string) (MusicProgram, error) {
	var m MusicProgram

	// create a new reader
	rd := r.New(r.NoLogger(), r.NoteOn(m.noteOn))

	// read the midi file
	err := r.ReadSMFFile(rd, file)

	return m, err
}

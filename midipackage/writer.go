package midipackage

import (
	"os"

	w "gitlab.com/gomidi/midi/writer"
)

// Write a Midi note to a SMF (Standard Midi File)
// delta is the length of the note (delta time between the **ON** and the **OFF** triggers)
// vel is the velocity in wich the note should be played
func (m MidiNote) writeNote(wr *w.SMF, delta uint32, vel uint8) {
	key := uint8(m)

	w.NoteOn(wr, key, vel)
	wr.SetDelta(delta)
	w.NoteOff(wr, key)
	wr.SetDelta(delta)
}

// Write a musicProgram to a SMF (Standard Midi File)
// All notes are written to channel 0
func (m *MusicProgram) Write(file string, delta uint32) error {
	// create file
	os.Remove(file)
	_, err := os.Create(file)
	if err != nil {
		return err
	}

	// Write all notes from the program to a SMF
	err = w.WriteSMF(file, 1, func(wr *w.SMF) error {

		wr.SetChannel(0)

		// loop through each note and write it with a:
		// velocity of 100
		// length of 120
		for _, note := range *m {
			note.writeNote(wr, delta, 100)
		}

		w.EndOfTrack(wr)

		return nil
	})
	return err
}

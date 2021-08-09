package midipackage

// Note type
// -> Represents a musical note
// -> Index is the number of the note on the 12th notes
//		-> C is 0; C# is 1; D is 2; ... ; B is 11
// -> Octave is the octave where the note is located
// -> Velocity is the intensity the note is played [0, 127]
// -> LenDivider is the number of this note that would fit into one beat
//		-> 4 is 1/4 beat; 10 is 1/10 beat ...
type Note struct {
	Index      int
	Octave     int
	Velocity   int
	LenDivider int
}

type ScaleType int

type Scale [7]Note

// Generate a new Note object based on the 3 numbers provided
func NewNote(note, vel, length int) Note {

	// velocity must be in [0, 127]
	vel = putInRange(vel, 0, 127)

	// length must be a natural (N) number
	length = natural(length)

	// get the Index
	scaleIndex := mod(note, numberOfNotes)

	// get the octave
	octave := ((note - scaleIndex) / numberOfNotes) + baseOctave

	return Note{Index: scaleIndex,
		Octave:     octave,
		Velocity:   vel,
		LenDivider: length}
}

package midipackage

// A scale is a slice of 8 Midi notes
// ex: CM (C major)
// C, D, E, F, G, A, B, C
type Scale [8]MidiNote

// A Midi note is an uint8
// it represents the index of a note on the keyboard
// C-1 --> 0
// C3 --> 48
type MidiNote uint8

// A music program is a slice of midi notes that constitutes a song
type MusicProgram []MidiNote

// Convert a string to a Scale
//
// M --> major
// m --> minor
// # --> sharp
// b --> flat
//
// C == CM == C major
// c == Cm == C minor
func NewScale(scale string) (Scale, error) {
	switch scale {
	case "CM", "C":
		return CM, nil
	case "C#M", "C#", "DbM", "Db":
		return CxM, nil
	case "DM", "D":
		return DM, nil
	case "D#M", "D#", "EbM", "Eb":
		return DxM, nil
	case "EM", "E":
		return EM, nil
	case "FM", "F":
		return FM, nil
	case "F#M", "F#", "GbM", "Gb":
		return FxM, nil
	case "GM", "G":
		return GM, nil
	case "G#M", "G#", "AbM", "Ab":
		return GxM, nil
	case "AM", "A":
		return AM, nil
	case "A#M", "A#", "BbM", "Bb":
		return AxM, nil
	case "BM", "B":
		return BM, nil
	case "Cm", "c":
		return Cm, nil
	case "C#m", "c#", "Dbm", "db":
		return Cxm, nil
	case "Dm", "d":
		return Dm, nil
	case "D#m", "d#", "Ebm", "eb":
		return Dxm, nil
	case "Em", "e":
		return Em, nil
	case "Fm", "f":
		return Fm, nil
	case "F#m", "f#", "Gbm", "gb":
		return FxM, nil
	case "Gm", "g":
		return Gm, nil
	case "G#m", "g#", "Abm", "ab":
		return Gxm, nil
	case "Am", "a":
		return Am, nil
	case "A#m", "a#", "Bbm", "bb":
		return Axm, nil
	case "Bm", "b":
		return Bm, nil
	default:
		return CM, ScaleError
	}
}

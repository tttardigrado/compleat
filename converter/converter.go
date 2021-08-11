package converter

import (
	bf "compleat/brainfuck"
	m "compleat/midipackage"
)

// Process a brainfuck program into a MusicProgram
// the notes contained in the outputed MusicProgram are dependent on the scale
// for more info on this see the bfToMidi function
func ProcessBF(bfProgram []rune, scale m.Scale) (notes m.MusicProgram) {
	// loop through every character
	for i := 0; i < len(bfProgram); i++ {
		switch bfProgram[i] {
		case '#':
			// ignore # comments
			// don't process any character untill the end of the line
			for i < len(bfProgram)-1 {
				i++
				if bfProgram[i] == '\n' {
					break
				}
			}
			break

		case '-': // 1st note on the scale
			notes = append(notes, scale[0])
			break

		case '+': // 2nd note on the scale
			notes = append(notes, scale[1])
			break

		case '<': // 3rd note on the scale
			notes = append(notes, scale[2])
			break

		case '>': // 4th note on the scale
			notes = append(notes, scale[3])
			break

		case '[': // 5th note on the scale
			notes = append(notes, scale[4])
			break

		case ']': // 6th note on the scale
			notes = append(notes, scale[5])
			break

		case '.': // 7th note on the scale
			notes = append(notes, scale[6])
			break

		case ',': // 8th note on the scale
			notes = append(notes, scale[7])
			break

		default:
			// ignore unrecognized characters
			break
		}
	}
	return
}

// read a brainfuck file and convert that program into midi
// the notes outputed will depend on the provided scale
//
// ex:
// | Operator | CM | CxM | ... |
// |:---------|:--:|:---:|:---:|
// | -        | C3 | Cx3 |     |
// | +        | D3 | Dx3 |     |
// | <        | E3 | F3  |     |
// | >        | F3 | Fx3 |     |
// | [        | G3 | Gx3 |     |
// | ]        | A3 | Ax3 |     |
// | .        | B3 | C4  |     |
// | ,        | C4 | Cx4 |     |
//
func BfToMidi(file string, scale m.Scale) (m.MusicProgram, error) {
	// read the brainfuck file
	// default extension for brainfuck files is either .bf or .b
	// bfProgram is a slice of runes containing all the characters in the file
	bfProgram, err := bf.Read(file)

	if err != nil {
		return []m.MidiNote{}, err
	}

	// Process the brainfuck program by converting each operator into a note
	notes := ProcessBF(bfProgram, scale)

	return notes, nil
}

// Process a MusicProgram program into a Brainfuck slice of runes
// the way this notes are converted to bf operators is dependent on the given scale
// for more info on this see the MidiToBF function
func ProcessMidi(midi m.MusicProgram, scale m.Scale) (bfProgram []rune) {
	// loop through every note
	for _, note := range midi {
		switch note {
		case scale[0]: // add -
			bfProgram = append(bfProgram, '-')
			break

		case scale[1]: // add +
			bfProgram = append(bfProgram, '+')
			break

		case scale[2]: // add <
			bfProgram = append(bfProgram, '<')
			break

		case scale[3]: // add >
			bfProgram = append(bfProgram, '>')
			break

		case scale[4]: // add [
			bfProgram = append(bfProgram, '[')
			break

		case scale[5]: // add ]
			bfProgram = append(bfProgram, ']')
			break

		case scale[6]: // add .
			bfProgram = append(bfProgram, '.')
			break

		case scale[7]: // add ,
			bfProgram = append(bfProgram, ',')
			break

		default:
			// ignore other notes
			break
		}
	}
	return
}

// read a midi file and convert the notes into brainfuck
// the way the notes will be processed is dependent on the given scale
//
// ex:
// | Operator | CM | CxM | ... |
// |:---------|:--:|:---:|:---:|
// | -        | C3 | Cx3 |     |
// | +        | D3 | Dx3 |     |
// | <        | E3 | F3  |     |
// | >        | F3 | Fx3 |     |
// | [        | G3 | Gx3 |     |
// | ]        | A3 | Ax3 |     |
// | .        | B3 | C4  |     |
// | ,        | C4 | Cx4 |     |
//
func MidiToBF(file string, scale m.Scale) ([]rune, error) {
	// read the midi file
	midiFile, err := m.Read(file)

	if err != nil {
		return []rune{}, err
	}

	// Process the midi file by turning every note on the scale
	// into it's correspondent brainfuck operator
	bfProgram := ProcessMidi(midiFile, scale)

	return bfProgram, nil
}

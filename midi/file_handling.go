package midipackage

import "io/ioutil"

// Read a file
// it returns the contents of the file as a slice of runes []rune
// the error returned is relative to trying to read the file
func loadMusicFileContents(file string) (string, error) {
	content, err := ioutil.ReadFile(file)

	// convert []byte -> string
	return string(content), err
}

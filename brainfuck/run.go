package brainfuck

import (
	"fmt"

	"io/ioutil"
)

// Run a slice of runes as a brainfuck commands
// to use a string convert it to a slice of runes
// with: []rune(stringYouWantToUse)
func RunBF(commands []rune) (Program, error) {
	program := NewProgram(commands)

	// if an error appens in the program
	// it will be returned from this function
	err := program.Run()

	fmt.Print("\n")
	return program, err
}

// Read a file
// it returns the contents of the file as a slice of runes []rune
// the error returned is relative to trying to read the file
func Read(file string) ([]rune, error) {
	content, err := ioutil.ReadFile(file)

	// convert []byte -> string -> []rune
	return []rune(string(content)), err
}

// Run a brainfuck program from a file.
func RunFile(file string) (Program, error) {
	content, err := Read(file)

	if err != nil {
		return NewProgram([]rune{}), err
	}

	program, err2 := RunBF(content)

	return program, err2
}

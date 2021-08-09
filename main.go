package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadMusicFile(filePath string) ([]string, error) {
	// slice containing all lines on the file
	lines := []string{}

	// Open and read file
	file, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer file.Close()

	// append all lines to the lines slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, fmt.Sprintf(scanner.Text()))
	}

	// check for errors
	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}

func RuneToBrainFuckOperator(char rune) rune {
	switch char {
	case '0':
		return ','
	case '1':
		return '-'
	case '2':
		return '+'
	case '3':
		return '<'
	case '4':
		return '>'
	case '5':
		return '.'
	case '6':
		return '['
	case '7':
		return ']'
	default:
		return ' '
	}
}

func LinesToBrainFuck(lines []string) (program []rune) {
	for _, line := range lines {
		for _, char := range line {
			operator := RuneToBrainFuckOperator(char)
			if operator != ' ' {
				program = append(program, operator)
			}
		}
	}
	return
}

func main() {
	lines, _ := LoadMusicFile("ideas.txt")
	fmt.Println(lines)
}

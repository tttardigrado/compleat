package brainfuck

import "fmt"

// integer type that represent a single brainfuck memory cell
type Cell int

// slice type that represent a collection of all brainfuck memory cells
type CellList []Cell

// Create a CellList object from a []int{} object (integer slice)
// All int in cellsSlice will be converted to Cell
func cellsFromIntSlice(cellsSlice []int) (c CellList) {
	for i := range cellsSlice {
		c = append(c, Cell(cellsSlice[i]))
	}

	return
}

// Increment a brainfuck memory cell
func (c *Cell) Increment() {
	// Maximum value of 255
	if *c >= CellMax {
		// wrap to 0
		*c = CellMin
	} else {
		// just increment
		*c += 1
	}
}

// Decrement a brainfuck memory cell
func (c *Cell) Decrement() {
	// Minimum value of 0
	if *c <= CellMin {
		// wrap to 255
		*c = CellMax
	} else {
		// just decrement
		*c -= 1
	}
}

// Print the content of the current cell
// The content will be converted from an integer to a rune
// ASCII: 65 will be printed as A
func (c *Cell) Print() {
	fmt.Printf("%c", rune(int(*c)))
}

// Set the content of a cell based on the input given by the user
func (c *Cell) SetFromInput() {
	// ask the user for a value
	var str string
	fmt.Scanf("%s", &str)

	// convert to a cell
	if str != "" {
		*c = Cell(int(str[0]))
	} else {
		*c = NullCell
	}
}

// Check if the cell is Null
// Cell(0) / NullCell is the null value for cell
func (c *Cell) IsNull() bool {
	return *c == NullCell
}

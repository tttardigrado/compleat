package brainfuck

import "errors"

// Errors
// Error when the cursor surpasses the length
var CursorError = errors.New("The value of the cursor is larger than the length of the memory's cells")

// Error when the cursor is moved to -1
var MoveDownError = errors.New("The cursor was moved to -1. The cursor has to be within [0, length[")

// Error when the cursor is moved to length
var MoveUpError = errors.New("The cursor was moved to more than it should. The cursor has to be within [0, length[")

// Error when the cursor is not on the right bounds
var ProgramBoundsError = errors.New("The cursor is either negative or larger than the length of the program")

// Constants
// Maximum value of a Cell
const CellMax = Cell(255)

// Minimum value of a Cell
const CellMin = Cell(0)

// Null value of a Cell
const NullCell = Cell(0)

package brainfuck

// struct type that represents the Memory tape
// of the brainfuck program
// made of:
//	cells (CellList) -> A list of all the cells
//  cursor (int) -> index of the current cell
type Memory struct {
	cells  CellList
	cursor int
}

// Generate a new Memory object
// cells is a slice of integers that will be converted to cells
func NewMemory(cells []int, cursor int) Memory {
	return Memory{cells: cellsFromIntSlice(cells), cursor: cursor}
}

// Add a new cell to the Memory
func (m *Memory) NewCell() {
	// Append NullCell
	m.cells = append(m.cells, NullCell)
}

// Get a pointer to the current cell
func (m *Memory) CurrentCell() *Cell {
	if m.cursor < len(m.cells) {
		return &m.cells[m.cursor]
	} else {
		return &m.cells[len(m.cells)-1]
	}
}

// Move the cursor of the Memory tape up
func (m *Memory) MoveUp() {
	if m.cursor < len(m.cells)-1 {
		m.cursor += 1
	} else {
		// if the cursor is at the edge of the tape
		// add a new cell and move to it
		m.NewCell()
		m.cursor += 1
	}
}

// Move the cursor of the Memory tape down
// returns an error if the cursor is a 0
func (m *Memory) MoveDown() error {
	if m.cursor == 0 {
		// cursor can't be negative
		return MoveDownError
	}

	// decrement
	m.cursor -= 1
	return nil
}

// Increment the value of the current cell
func (m *Memory) Increment() {
	current := m.CurrentCell()

	current.Increment()
}

// Decrement the value of the current cell
func (m *Memory) Decrement() {
	current := m.CurrentCell()

	current.Decrement()
}

// Print the content of the current cell
// The content will be converted from an integer to a rune
// ASCII: 65 will be printed as A
func (m *Memory) Print() {
	current := m.CurrentCell()

	current.Print()
}

// Set the content of the current cell
// based on the input given by the user
func (m *Memory) Ask() {
	current := m.CurrentCell()

	current.SetFromInput()
}

// Check if the cell is Null
// Cell(0) / NullCell is the null value for cell
func (m *Memory) IsCurrentNull() bool {
	return m.CurrentCell().IsNull()
}

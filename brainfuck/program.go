package brainfuck

// struct type for the brainfuck program
// made of:
//	Commands ([]rune) -> A list of all the commands in the program
//  Cursor (int) -> index of the current command
//	Mem (Memory) -> Memory object, where the tape memory of the program is stored
type Program struct {
	Commands []rune
	Cursor   int
	Mem      Memory
}

// Get the current command of the program and process it
func (p *Program) InterpretCommand() error {
	command, err := p.GetCurrentCommand()

	if err != nil {
		return err
	}

	// process command
	switch command {
	case '+':
		// C ++*ptr;
		p.Mem.Increment()

	case '-':
		// C --*ptr;
		p.Mem.Decrement()

	case '>':
		// C ++ptr;
		p.Mem.MoveUp()

	case '<':
		// C --ptr;
		// this can result in error if the cursor is moved to -1
		err := p.Mem.MoveDown()
		return err

	case '.':
		// C putchar(*ptr);
		p.Mem.Print()

	case ',':
		// C *ptr = getchar();
		p.Mem.Ask()

	case '[':
		// C while (*ptr){
		p.BeginLoop()

	case ']':
		// C }
		p.EndLoop()

	default:
		// don't process
		break
	}
	return nil
}

// Get the length of the commands of the program
func (p *Program) Length() int {
	return len(p.Commands)
}

// Get the current command based on the cursor position
// returns it's value and an error
func (p *Program) GetCurrentCommand() (rune, error) {
	length := p.Length()

	// check if the cursor is in bounds
	isInBounds := p.Cursor >= 0 && p.Cursor < length

	if isInBounds {
		// return the command and no error
		return p.Commands[p.Cursor], nil
	}

	// return empty command and an error
	return ' ', ProgramBoundsError
}

// Increment the program's Cursor by 1
func (p *Program) IncrementCursor() {
	p.Cursor += 1
}

// Decrement the program's Cursor by 1
func (p *Program) DecrementCursor() {
	p.Cursor -= 1
}

// Run a program, by iterating over it's commands
// and processing each of them.
//
// if an error occurs during the excution
// it will return the apropriate error
func (p *Program) Run() error {
	// iterate over every command
	for p.Cursor < p.Length() {
		err := p.InterpretCommand()

		if err != nil {
			return err
		}

		p.IncrementCursor()
	}

	return nil
}

// Generate a new default Program Object from a slice of runes (commands)
// cursor starts at zero
// mem starts as {[0], 0}
func NewProgram(commands []rune) Program {

	return Program{
		Commands: commands,
		Cursor:   0,
		Mem:      NewMemory([]int{0}, 0)}
}

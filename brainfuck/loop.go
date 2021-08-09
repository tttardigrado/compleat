package brainfuck

// Start a brainfuck loop
// The loop starts on the '[' command
func (p *Program) BeginLoop() error {
	// When the cell is null don't start the loop and
	// go to the matching ']'
	if p.Mem.IsCurrentNull() {
		// number of opening square brackets '['
		counter := 0

		// move up so that the starting '[' is not considered
		p.IncrementCursor()

		for p.Cursor < p.Length() {
			current, err := p.GetCurrentCommand()

			if err != nil {
				return err
			}

			if current == ']' && counter == 0 {
				// matching character -> quit loop
				break
			} else if current == '[' {
				// new ']' character -> decrement counter
				counter += 1
			} else if current == ']' {
				// new ']' character -> decrement counter
				counter -= 1
			}

			// next character
			p.IncrementCursor()
		}

	}

	// continue the loop
	return nil
}

// End a brainfuck loop
// The loop starts on the ']' command
func (p *Program) EndLoop() error {
	// When the cell is not null go back to
	// the beggining of the loop,
	// the matching '['
	if !p.Mem.IsCurrentNull() {
		// number of closing square brackets ']'
		counter := 0

		// move down so that the starting ']' is not considered
		p.DecrementCursor()

		for p.Cursor >= 0 {
			current, err := p.GetCurrentCommand()

			if err != nil {
				return err
			}

			if current == '[' && counter == 0 {
				// matching character -> quit loop
				break
			} else if current == ']' {
				// new ']' character -> increment counter
				counter += 1
			} else if current == '[' {
				// new '[' character -> decrement counter
				counter -= 1
			}

			// next character
			p.DecrementCursor()
		}

	}

	// exit loop and
	// continue the program
	return nil
}

package midipackage

// Modulus/Rest function
// can properly handle negative numbers
func mod(a, b int) int {
	return (a%b + b) % b
}

// Absolute value of a integer
// abs(x) == |x|
// abs(-10) == abs(10) == 10
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// integer to Natural number
// natural(-10) == natural(10) == 10
// natural(0) == 1 (because 0 is not a natural number)
func natural(n int) int {
	if n > 0 {
		return n
	} else if n < 0 {
		return -n
	}
	return 1
}

// Put a number (n) in the specified range
func putInRange(n, min, max int) int {
	// max needs to be larger than min
	// this switches them if min is larger
	if max < min {
		tmp := max
		max = min
		min = tmp
	}

	if n > max {
		// more than maximum
		return max
	} else if n < min {
		// less than minimum
		return min
	}
	// already in range
	return n
}

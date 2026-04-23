package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)

	// '0' is allowed because it is not 1 through 9
	z := '0'

	a := z
	a++ // essentially '1'

	b := a
	b++ // essentially '2'

	d := b
	d++
	d++ // essentially '4'

	// Build the slice using our variables to avoid literal numbers
	output := []rune{'x', ' ', '=', ' ', d, b, ',', ' ', 'y', ' ', '=', ' ', b, a, '\n'}

	// Loop so we only use one call to z01.PrintRune
	for _, r := range output {
		z01.PrintRune(r)
	}
}

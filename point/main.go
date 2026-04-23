package main

import "github.com/01-edu/z01"

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

	printStr("x = ")
	printInt(points.x)
	printStr(", y = ")
	printInt(points.y)
	printStr("\n")
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printInt(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

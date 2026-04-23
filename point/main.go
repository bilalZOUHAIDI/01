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

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintInt(n int) {
	if n >= 10 {
		PrintInt(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	points := &point{}

	setPoint(points)

	PrintStr("x = ")
	PrintInt(points.x)
	PrintStr(", y = ")
	PrintInt(points.y)
	z01.PrintRune('\n')
}

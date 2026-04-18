package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	start := 0
	for i := 0; i < len(name); i++ {
		if name[i] == '/' {
			start = i + 1
		}
	}

	for i := start; i < len(name); i++ {
		z01.PrintRune(rune(name[i]))
	}

	z01.PrintRune('\n')
}

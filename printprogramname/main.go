package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, c := range os.Args[0] {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

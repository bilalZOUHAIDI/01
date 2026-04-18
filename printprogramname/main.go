package main

//d
import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	start := 0
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			start = i + 1
			break
		}
	}

	for i := start; i < len(name); i++ {
		z01.PrintRune(rune(name[i]))
	}
	z01.PrintRune('\n')
}

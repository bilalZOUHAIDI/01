package main

//osa
import (
	"os"
	"path/filepath"

	"github.com/01-edu/z01"
)

func main() {
	name := filepath.Base(os.Args[0])

	for _, c := range name {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

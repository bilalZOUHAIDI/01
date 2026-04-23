package main

import (
	"bufio"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		printStr(scanner.Text())
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printFile(os.Stdin)
		return
	}

	for _, filename := range args {
		file, err := os.Open(filename)
		if err != nil {
			printStr("ERROR: open " + filename + ": no such file or directory")
			os.Exit(1)
		}
		defer file.Close()
		printFile(file)
	}
}

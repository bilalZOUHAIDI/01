package main

import (
	"bufio"
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printError(msg string) {
	for _, r := range msg {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printFile(name string) bool {
	file, err := os.Open(name)
	if err != nil {
		printError("ERROR: open " + name + ": " + err.Error())
		return false
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		r, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			printError("ERROR: " + err.Error())
			return false
		}
		z01.PrintRune(rune(r))
	}

	return true
}

func printStdin() {
	reader := bufio.NewReader(os.Stdin)

	for {
		r, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			printError("ERROR: " + err.Error())
			os.Exit(1)
		}
		z01.PrintRune(rune(r))
	}
}

func main() {
	args := os.Args[1:]

	// no args → stdin mode
	if len(args) == 0 {
		printStdin()
		return
	}

	exitCode := 0

	for _, file := range args {
		ok := printFile(file)
		if !ok {
			exitCode = 1
		}
	}

	if exitCode == 1 {
		os.Exit(1)
	}
}

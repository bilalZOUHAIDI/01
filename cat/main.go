package main

import (
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
		printError("ERROR: " + err.Error())
		return false
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		printError("ERROR: " + err.Error())
		return false
	}

	return true
}

func printStdin() {
	_, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		printError("ERROR: " + err.Error())
		os.Exit(1)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printStdin()
		return
	}

	exitCode := 0

	for _, file := range args {
		if !printFile(file) {
			exitCode = 1
		}
	}

	if exitCode == 1 {
		os.Exit(1)
	}
}

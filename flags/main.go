package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func getValue(arg string, prefix string) string {
	res := ""
	for i := len(prefix); i < len(arg); i++ {
		res += string(arg[i])
	}
	return res
}

func sortASCII(s string) string {
	r := []rune(s)

	for i := 0; i < len(r)-1; i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return string(r)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	insert := ""
	order := false
	mainStr := ""

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if hasPrefix(arg, "--insert=") {
			insert = getValue(arg, "--insert=")
		} else if hasPrefix(arg, "-i=") {
			insert = getValue(arg, "-i=")
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			mainStr = arg
		}
	}

	result := mainStr

	if insert != "" {
		result = result + insert
	}

	if order {
		result = sortASCII(result)
	}

	fmt.Println(result)
}

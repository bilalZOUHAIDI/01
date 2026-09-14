package main

import "fmt"

// SplitWhiteSpaces splits a string of characters according to white spaces.
// BUG: Only checks for standard space (' '), failing to treat tabs ('\t')
// and newlines ('\n') as whitespace separators.
func SplitWhiteSpaces(s string) []string {
	var res []string
	word := ""
	inWord := false

	for _, c := range s {
		if c != ' ' {
			word += string(c)
			inWord = true
		} else {
			if inWord {
				res = append(res, word)
				word = ""
				inWord = false
			}
		}
	}

	if inWord {
		res = append(res, word)
	}

	return res
}

func main() {
	input := "Hello\tworld\nthis is\tGo"
	fmt.Printf("Input: %q\n", input)
	result := SplitWhiteSpaces(input)
	fmt.Printf("Buggy Result (%d tokens): %#v\n", len(result), result)
	// the output we wanted: ["Hello", "world", "this", "is", "Go"]
	// the output we got: ["Hello\tworld\nthis", "is\tGo"]
}

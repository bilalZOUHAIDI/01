package main

import "fmt"

func isWhiteSpace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

// SplitWhiteSpaces splits a string of characters according to white spaces
// (' ', '\t', and '\n').
func SplitWhiteSpaces(s string) []string {
	var res []string
	word := ""
	inWord := false

	for _, c := range s {
		if !isWhiteSpace(c) {
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
	fmt.Printf("Fixed Result (%d tokens): %#v\n", len(result), result)
	// Output: ["Hello", "world", "this", "is", "Go"]
}

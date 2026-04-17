package main

import (
	"fmt"
)

func Compare(a, b string) int {
	i := 0

	for i < len(a) && i < len(b) {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
		i++
	}

	// If all characters so far are equal, shorter string is smaller
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}

	return 0
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}

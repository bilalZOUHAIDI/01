package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	upper := false

	// Check for flag
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	result := ""

	for _, arg := range args {
		n, err := strconv.Atoi(arg)

		if err != nil || n < 1 || n > 26 {
			result += " "
			continue
		}

		if upper {
			result += string(rune('A' + n - 1))
		} else {
			result += string(rune('a' + n - 1))
		}
	}

	fmt.Println(result)
}

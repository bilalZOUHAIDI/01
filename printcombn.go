package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)

	// Initialize first combination
	for i := 0; i < n; i++ {
		comb[i] = i
	}

	for {
		printComb(comb, n)

		// Find position to increment
		i := n - 1
		for i >= 0 && comb[i] == 10-n+i {
			i--
		}

		if i < 0 {
			break
		}

		comb[i]++

		// Reset following digits
		for j := i + 1; j < n; j++ {
			comb[j] = comb[j-1] + 1
		}
	}
}

func printComb(comb []int, n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(rune(comb[i] + '0'))
	}

	// Check if it's the last combination
	if comb[0] != 10-n {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}


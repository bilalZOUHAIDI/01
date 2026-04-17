package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {

	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	baseLen := len(base)

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	var digits [64]int
	i := 0
	for nbr > 0 {
		digits[i] = nbr % baseLen
		nbr /= baseLen
		i++
	}

	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(rune(base[digits[j]]))
	}
}

package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNbrRecursive(n)
}

func PrintNbrRecursive(n int) {
	if n == 0 {
		return
	}
	PrintNbrRecursive(n / 10)
	d := n % 10
	if d < 0 {
		d = -d
	}
	for i, r := range "0123456789" {
		if i == d {
			z01.PrintRune(r)
		}
	}
}

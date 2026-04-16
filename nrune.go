package piscine

func NRune(s string, n int) rune {
	for i, r := range s {
		if i == n-1 {
			return r
		}
	}
	return 0
}

package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0 // invalid character
		}
		result = result*10 + int(ch-'0')
	}

	return result
}

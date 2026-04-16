package piscine

func ToUpper(s string) string {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r -= 32
		}
	}
	return s
}

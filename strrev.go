package piscine

// This is a single-line comment

func StrRev(s string) string {
	B := ""
	for i := len(s) - 1; i >= 0; i-- {
		B = B + string(s[i])
	}
	return B
}

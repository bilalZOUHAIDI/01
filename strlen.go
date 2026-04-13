package piscine

func Strlen(s string) int {
	len := 0
	for range s {
		len = len + 1
	}
	return len
}

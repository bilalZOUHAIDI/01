package piscine

func Join(elems []string, sep string) string {
	var result string
	for i, c := range elems {
		if i == len(elems)-1 {
			result += c
		} else {
			result += c + sep
		}
	}
	return result
}

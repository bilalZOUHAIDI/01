package piscine

func BasicJoin(elems []string) string {
	var result string
	for _, c := range elems {
		result += c
	}
	return result
}

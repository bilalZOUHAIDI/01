package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	word := ""
	inWord := false

	for _, c := range s {
		if c != ' ' {
			word += string(c)
			inWord = true
		} else {
			if inWord {
				res = append(res, word)
				word = ""
				inWord = false
			}
		}
	}

	if inWord {
		res = append(res, word)
	}

	return res
}

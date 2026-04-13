package piscine

import "strconv"

func BasicAtoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	} else {
		return result
	}
}

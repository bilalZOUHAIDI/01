package piscine

import "strconv"

// fo
func BasicAtoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		return 0
	} else {
		return res
	}

}

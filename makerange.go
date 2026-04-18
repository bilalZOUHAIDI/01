package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int{}
	}

	res := make([]int, max-min)

	for i := min; i < max; i++ {
		res[i-min] = i
	}

	return res
}

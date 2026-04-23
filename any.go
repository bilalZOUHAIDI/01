package piscine

func Any(f func(int) bool, a []int) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}

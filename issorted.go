package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}

	i := 0
	for i < len(a)-1 && f(a[i], a[i+1]) == 0 {
		i++
	}

	if i == len(a)-1 {
		return true
	}

	ascending := f(a[i], a[i+1]) < 0

	for i < len(a)-1 {
		if ascending {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		} else {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
		i++
	}

	return true
}

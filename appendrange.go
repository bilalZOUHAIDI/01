

func AppendRange(start, end int) []int {
	var tab []int

	for i := start; i <= end; i++ {
		tab = append(tab, i)
	}

	return tab
}

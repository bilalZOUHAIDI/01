package piscine

func Compact(ptr *[]string) int {
	result := make([]string, 0, len(*ptr))

	for _, v := range *ptr {
		if v != "" {
			result = append(result, v)
		}
	}

	*ptr = result
	return len(result)
}

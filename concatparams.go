package piscine

func ConcatParams(args []string) string {
	var res string
	for _, arg := range args {
		res += arg
	}
	return res
}

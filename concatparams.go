package piscine

func ConcatParams(args []string) string {
	var res string
	for i, arg := range args {
		res += arg
		if i != len(args)-1 {
			res += "\n"
		}
	}
	return res
}

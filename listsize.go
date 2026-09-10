package piscine

func ListSize(l *List) int {
	if l == nil {
		return 0
	}
	count := 0
	for it := l.Head; it != nil; it = it.Next {
		count++
	}
	return count
}

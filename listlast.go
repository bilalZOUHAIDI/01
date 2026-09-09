package piscine

func ListLast(l *List) interface{} {
	if l == nil || l.Head == nil {
		return nil
	}
	it := l.Head
	for it.Next != nil {
		it = it.Next
	}
	return it.Data
}

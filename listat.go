package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	for curr := l; curr != nil; curr = curr.Next {
		if pos == 0 {
			return curr
		}
		pos--
	}
	return nil
}

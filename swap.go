package piscine

func Swap(a *int, b *int) {
	d := *a
	*a = *b
	*b = d
}

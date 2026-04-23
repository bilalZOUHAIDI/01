package piscine

import "fmt"

func PrintNbr(n int) {
	fmt.Print(n)
}

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}

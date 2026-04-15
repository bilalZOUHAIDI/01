package main

import "fmt"

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 {
		return 0
	} else {
		for i := 2; i <= nb; i++ {
			result *= i
		}
	}
	return result
}

func main() {
	fmt.Println(IterativeFactorial(5))
}

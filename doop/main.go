package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <number1> <operator> <number2>")
		return
	}

	n1, err1 := strconv.Atoi(os.Args[1])
	op := os.Args[2]
	n2, err2 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid numbers")
		return
	}

	var result int
	switch op {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		if n2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = n1 / n2
	case "%":
		if n2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = n1 % n2
	default:
		fmt.Println("Error: Invalid operator")
		return
	}

	fmt.Println(result)
}

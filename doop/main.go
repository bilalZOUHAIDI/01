package main

import "os"

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok1 := Atoi(os.Args[1])
	b, ok2 := Atoi(os.Args[3])
	op := os.Args[2]

	if !ok1 || !ok2 || len(op) != 1 {
		return
	}

	var result int

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			PrintStr("No division by 0\n")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			PrintStr("No modulo by 0\n")
			return
		}
		result = a % b
	default:
		return
	}

	PrintNbr(result)
	PrintStr("\n")
}

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	sign := 1
	i := 0

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	result := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := int(s[i] - '0')

		if result > (9223372036854775807-digit)/10 {
			return 0, false
		}

		result = result*10 + digit
	}

	return result * sign, true
}

func PrintStr(s string) {
	os.Stdout.Write([]byte(s))
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		PrintStr("-9223372036854775808")
		return
	}

	if n < 0 {
		PrintStr("-")
		n = -n
	}

	if n >= 10 {
		PrintNbr(n / 10)
	}

	digit := byte(n%10 + '0')
	os.Stdout.Write([]byte{digit})
}

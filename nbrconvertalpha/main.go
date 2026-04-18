package main

import "os"

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	n := 0
	for i := 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return n, true
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	out := make([]byte, 0, len(args))

	for _, arg := range args {
		n, ok := atoi(arg)

		if !ok || n < 1 || n > 26 {
			out = append(out, ' ')
		} else if upper {
			out = append(out, byte('A'+n-1))
		} else {
			out = append(out, byte('a'+n-1))
		}
	}

	if len(out) > 0 {
		out = append(out, '\n')
		os.Stdout.Write(out)
	}
}

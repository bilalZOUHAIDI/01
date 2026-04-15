package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int

	var solve func(int)
	solve = func(col int) {
		if col == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(board[i] + '0'))
			}
			z01.PrintRune('\n')
			return
		}

		for row := 1; row <= 8; row++ {
			ok := true
			for i := 0; i < col; i++ {
				if board[i] == row || (board[i]-row == col-i) || (row-board[i] == col-i) {
					ok = false
					break
				}
			}
			if ok {
				board[col] = row
				solve(col + 1) // recursion
			}
		}
	}

	solve(0)
}

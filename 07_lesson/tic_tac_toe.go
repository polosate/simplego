package _7_lesson

import "errors"

// 0 - 0
// 1 - x
// 2 - empty
func isWon(board [][]int) (int, error) {
	if len(board[0]) != len(board) {
		return 2, errors.New("")
	}
	n := len(board[0])
	var row, col int

	// check rows
	for row = 0; row < n; row++ {
		if board[row][0] != 2 {
			for col = 1; col < n; col++ {
				if board[row][col] != board[row][col-1] {
					break
				}
			}
			if col == n {
				return board[row][0], nil
			}
		}
	}

	// check cols
	for col = 0; col < n; col++ {
		if board[0][col] != 2 {
			for row = 1; row < n; row++ {
				if board[row][col] != board[row-1][col] {
					break
				}
			}
			if row == n {
				return board[0][row], nil
			}
		}
	}

	// check diags
	// TODO

	return 2, nil
}

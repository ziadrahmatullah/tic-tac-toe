// Name : Muhammad Ziad Rahmatullah
// Date : Okt 12, 20:24

package main

import (
	"fmt"
	"regexp"
	"strings"
)

var winPatterns = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
	{0, 4, 8}, {2, 4, 6},            // Diagonals
}

func checkInvalidGameBoard(board string) bool{
	if matched, _ := regexp.MatchString("^[XO\\-]{9}$", board); !matched {
		return false
	}

	xCount := strings.Count(board, "X")
	oCount := strings.Count(board, "O")
	if !(xCount == oCount+1) && !(xCount+1 == oCount) {
		return false
	}
	return true
}

func checkTicTacToeState(board string) string {
	var xWins, oWins bool
	for _, pattern := range winPatterns {
		if board[pattern[0]] == board[pattern[1]] && board[pattern[1]] == board[pattern[2]] {
			if board[pattern[0]] == 'X' {
				xWins = true
			} else if board[pattern[0]] == 'O' {
				oWins = true
			}
		}
	}
	if xWins && oWins || !checkInvalidGameBoard(board) {
		return "Invalid game board"
	}
	if xWins {
		return "X wins!"
	}
	if oWins {
		return "O wins!"
	}
	if strings.Contains(board, "-") {
		return "Game still in progress!"
	}
	return "It's a draw!"
}

func main() {
	var board string
	fmt.Print("Input board :")
	fmt.Scanln(&board)
	result := checkTicTacToeState(board)
	fmt.Println(result)
}

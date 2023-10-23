package game

import (
	"regexp"
	"strings"

	CustomError "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/tic-tac-toe/errors"
)

var winPatterns = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
	{0, 4, 8}, {2, 4, 6},            // Diagonals
}

func isInvalidGameBoard(board string) (bool, error){
	if matched, _ := regexp.MatchString("^[XO\\-]{9}$", board); !matched {
		return false, &CustomError.TicTacToeError{Msg: "board contain except X O or less than 9 char"}
	}

	xCount := strings.Count(board, "X")
	oCount := strings.Count(board, "O")
	if !(xCount == oCount+1) && !(xCount+1 == oCount) {
		return false, &CustomError.TicTacToeError{Msg:"x and o false count" }
	}
	return true, nil
}

func TicTacToe(board string) (output string, err error) {
	if ok, err:=isInvalidGameBoard(board); !ok {
		return "", err
	}
	if strings.Contains(board, "-") {
		return "", &CustomError.TicTacToeError{Msg: "game still in proggres"}
	}
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
	
	if xWins && oWins{
		return "", &CustomError.TicTacToeError{Msg: "invalid win state" }
	}
	if xWins {
		return "X wins!", nil
	}
	if oWins {
		return "O wins!", nil
	}
	return "It's a draw!", nil
}

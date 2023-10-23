package main

import (
	"fmt"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/tic-tac-toe/game"
	CustomError "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/tic-tac-toe/errors"
)

func main(){
	board := "XXOO"
	var TttErr *CustomError.TicTacToeError
	result, err :=game.TicTacToe(board)
	if err != nil{
		switch{
		case errors.As(err, &TttErr):
			fmt.Printf("%s have problem: %s\n", board, TttErr.Error())
		default:
			fmt.Printf("unexpected tictactoe error: %s", err)
		}
	}
	fmt.Println(result)
}
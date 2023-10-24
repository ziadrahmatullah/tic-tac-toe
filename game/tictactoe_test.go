package game_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/tic-tac-toe/errors"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/tic-tac-toe/game"
	"github.com/stretchr/testify/assert"
)

func TestTicTacToe(t *testing.T) {
	t.Run("should return 'X wins!' when input XOXXOOXXO", func(t *testing.T) {
		a := "XOXXOOXXO"
		expected := "X wins!"

		result, _ := game.TicTacToe(a)

		assert.Equal(t, expected, result)
	})

	t.Run("should return 'O wins!' when input XOOXOXOXO", func(t *testing.T) {
		a := "XOOXOXOXO"
		expected := "O wins!"

		result, _ := game.TicTacToe(a)

		assert.Equal(t, expected, result)
	})

	t.Run("should return 'It's a draw!' when input OXOXOXXOX", func(t *testing.T) {
		a := "OXOXOXXOX"
		expected := "It's a draw!"

		result, _ := game.TicTacToe(a)

		assert.Equal(t, expected, result)
	})

	t.Run("should return 'game still in proggres' when input XOXX--O--", func(t *testing.T) {
		a := "XOXX--O--"
		expected := &errors.TicTacToeError{Msg: "game still in proggres"}
		_, err:= game.TicTacToe(a)

		assert.EqualError(t, expected, err.Error())
	})

	t.Run("should return 'invalid win state' when input XXXOOOXXO", func(t *testing.T) {
		a := "XXXOOOXXO"
		expected := &errors.TicTacToeError{Msg: "invalid win state"}
		_, err:= game.TicTacToe(a)

		assert.EqualError(t, expected, err.Error())
	})

	t.Run("should return 'board contain except X O or less than 9 char' when input XXXOLLLXO", func(t *testing.T) {
		a := "XXXOLLLXO"
		expected := &errors.TicTacToeError{Msg: "board contain except X O or less than 9 char"}
		_, err:= game.TicTacToe(a)

		assert.EqualError(t, expected, err.Error())
	})

	t.Run("should return 'x and o false count' when input XXXXXXXXX", func(t *testing.T) {
		a := "XXXXXXXXX"
		expected := &errors.TicTacToeError{Msg: "x and o false count"}
		_, err:= game.TicTacToe(a)

		assert.EqualError(t, expected, err.Error())
	})
}
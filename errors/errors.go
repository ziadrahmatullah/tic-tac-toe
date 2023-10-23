package errors

type TicTacToeError struct {
    Msg  string
}

func (e *TicTacToeError) Error() string {
    return e.Msg
}

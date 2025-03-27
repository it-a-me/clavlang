package interpreter

import (
	"fmt"

	"github.com/it-a-me/clavlang/token"
)

type InterpreterError struct {
	message string
	token   token.Token
}

func newInterpreterError(message string, token token.Token) InterpreterError {
	return InterpreterError{
		message: message, token: token,
	}
}

func (i InterpreterError) Error() string {
	return fmt.Sprintf("Error on line %d: %s", i.token.Line, i.message)
}

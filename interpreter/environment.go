package interpreter

import (
	"errors"

	"github.com/it-a-me/clavlang/token"
	"github.com/it-a-me/clavlang/types"
)

type Environment struct {
	env map[string]types.ClavType
}

func NewEnvironment() Environment {
	return Environment{env: make(map[string]types.ClavType)}
}

func (e *Environment) Define(name string, value types.ClavType) {
	e.env[name] = value
}

func (e *Environment) Get(name token.Token) (types.ClavType, error) {
	if val, ok := e.env[name.Lexeme]; ok {
		return val, nil
	}

	return nil, errors.New("Undefined variable '" + name.Lexeme + "'")
}

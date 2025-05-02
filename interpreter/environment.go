package interpreter

import (
	"errors"
	"slices"

	"github.com/it-a-me/clavlang/token"
	"github.com/it-a-me/clavlang/types"
)

type Environment struct {
	env []map[string]types.ClavType
}

func NewEnvironment() Environment {
	env := []map[string]types.ClavType{}
	// global scope
	env = append(env, map[string]types.ClavType{})
	return Environment{env: env}
}

func (e *Environment) NewScope() {
	e.env = append(e.env, map[string]types.ClavType{})
}

func (e *Environment) EndScope() {
	e.env = e.env[0 : len(e.env)-1]
}

func (e *Environment) Define(name string, value types.ClavType) {
	e.env[len(e.env)-1][name] = value
}

func (e *Environment) Assign(name token.Token, value types.ClavType) (types.ClavType, error) {
	for _, env := range slices.Backward(e.env) {
		if _, ok := env[name.Lexeme]; ok {
			env[name.Lexeme] = value
			return value, nil
		}
	}

	return nil, errors.New("Undefined variable '" + name.Lexeme + "'")
}

func (e *Environment) Get(name token.Token) (types.ClavType, error) {
	for _, env := range slices.Backward(e.env) {
		if val, ok := env[name.Lexeme]; ok {
			return val, nil
		}
	}

	return nil, errors.New("Undefined variable '" + name.Lexeme + "'")
}

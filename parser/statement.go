package parser

import "github.com/it-a-me/clavlang/token"

type Stmt interface {
	stmt()
}

type Expression struct {
	Inner Expr
}

type Print struct {
	Inner Expr
}

type Var struct {
	Name        token.Token
	Initializer Expr
}

func (Expression) stmt() {}
func (Print) stmt()      {}
func (Var) stmt()        {}

package parser

import (
	"fmt"
	"reflect"

	"github.com/it-a-me/clavlang/token"
	"github.com/it-a-me/clavlang/types"
)

type Expr interface {
	expr()
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value types.ClavType
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

type Variable struct {
	Name token.Token
}

func LispStmt(stmt Stmt) string {
	switch s := stmt.(type) {
	case Print:
		return "(print " + LispExpr(s.Inner) + ")"
	case Expression:
		return LispExpr(s.Inner)
	}
	panic("Unreachable")
}

func LispExpr(expr Expr) string {
	if l, ok := expr.(Literal); ok {
		return fmt.Sprintf("%v", l.Value)
	}

	value := reflect.ValueOf(expr)
	s := "("
	for i := range value.NumField() {
		f := value.Field(i)
		if i != 0 {
			s += " "
		}
		switch v := f.Interface().(type) {
		case Expr:
			s += LispExpr(v)
		case token.Token:
			s += v.Lexeme
		default:
			s += f.String()
		}
	}
	s += ")"
	return s
}

func (Binary) expr()   {}
func (Grouping) expr() {}
func (Literal) expr()  {}
func (Unary) expr()    {}
func (Variable) expr() {}

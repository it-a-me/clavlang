package parser

import (
	"fmt"
	"reflect"

	"github.com/it-a-me/clavlang/token"
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
	Value any
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func LispExpr(expr Expr) string {
	if l, ok := expr.(Literal); ok {
		return fmt.Sprintf("%v", l.Value)
	}

	value := reflect.ValueOf(expr)
	s := "(" + value.Type().Name()
	for i := range value.NumField() {
		f := value.Field(i)
		switch v := f.Interface().(type) {
		case Expr:
			s += " " + LispExpr(v)
		case token.Token:
			s += " " + v.Lexeme
		default:
			s += " " + f.String()
		}
	}
	s += ")"
	return s
}

func (b Binary) expr()   {}
func (g Grouping) expr() {}
func (l Literal) expr()  {}
func (u Unary) expr()    {}

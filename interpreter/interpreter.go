package interpreter

import (
	"fmt"
	"reflect"

	"github.com/it-a-me/clavlang/parser"
	"github.com/it-a-me/clavlang/token"
	"github.com/it-a-me/clavlang/types"
)

func (i *Interpreter) Interpret(statements []parser.Stmt) error {
	for _, stmt := range statements {
		if err := i.execute(stmt); err != nil {
			return err
		}
	}
	return nil
}

type Interpreter struct {
	environment Environment
}

func NewInterpreter() Interpreter {
	return Interpreter{environment: NewEnvironment()}
}

func (i *Interpreter) execute(stmt parser.Stmt) error {
	switch s := stmt.(type) {
	case parser.Print:
		val, err := i.evaluate(s.Inner)
		if err != nil {
			return err
		}
		fmt.Println(val.String())
	case parser.Expression:
		_, err := i.evaluate(s.Inner)
		if err != nil {
			return err
		}
	case parser.Var:
		var value types.ClavType
		var err error
		if s.Initializer != nil {
			value, err = i.evaluate(s.Initializer)
			if err != nil {
				return err
			}
		}
		i.environment.Define(s.Name.Lexeme, value)
	}
	return nil
}

func (i *Interpreter) evaluate(expr parser.Expr) (types.ClavType, error) {
	switch e := expr.(type) {
	case parser.Literal:
		return i.evalutateLiteral(e), nil
	case parser.Grouping:
		return i.evalutateGrouping(e)
	case parser.Unary:
		return i.evalutateUnary(e)
	case parser.Binary:
		return i.evalutateBinary(e)
	case parser.Variable:
		return i.environment.Get(e.Name)
	}
	panic("Unreachable")
}

func (i *Interpreter) evalutateLiteral(expr parser.Literal) types.ClavType {
	return expr.Value
}

func (i *Interpreter) evalutateGrouping(expr parser.Grouping) (types.ClavType, error) {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) evalutateUnary(expr parser.Unary) (types.ClavType, error) {
	right, err := i.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Operator.Type {
	case token.Bang:
		old := right.(types.Boolean)
		return types.Boolean{Value: !old.Value}, nil
	case token.Minus:
		old := right.(types.Number)
		return types.Number{Value: -old.Value}, nil
	}
	panic("Unreachable")
}

func (i *Interpreter) evalutateBinary(expr parser.Binary) (types.ClavType, error) {
	left, err := i.evaluate(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := i.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Operator.Type {
	case token.Minus:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot subtract non-numeric type "+t, expr.Operator)
		}
		value := left.(types.Number).Value - right.(types.Number).Value
		return types.Number{Value: value}, nil
	case token.Slash:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot divide non-numeric type "+t, expr.Operator)
		}
		value := left.(types.Number).Value / right.(types.Number).Value
		return types.Number{Value: value}, nil
	case token.Star:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot multiply non-numeric type "+t, expr.Operator)
		}
		value := left.(types.Number).Value * right.(types.Number).Value
		return types.Number{Value: value}, nil
	case token.EqualEqual:
		eq, err := isEqual(expr.Operator, left, right)
		return types.Boolean{Value: eq}, err
	case token.BangEqual:
		eq, err := isEqual(expr.Operator, left, right)
		return types.Boolean{Value: !eq}, err
	case token.Greater:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		value := left.(types.Number).Value > right.(types.Number).Value
		return types.Boolean{Value: value}, nil
	case token.GreaterEqual:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		value := left.(types.Number).Value >= right.(types.Number).Value
		return types.Boolean{Value: value}, nil
	case token.Less:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		value := left.(types.Number).Value < right.(types.Number).Value
		return types.Boolean{Value: value}, nil
	case token.LessEqual:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		value := left.(types.Number).Value <= right.(types.Number).Value
		return types.Boolean{Value: value}, nil
	case token.Plus:
		if l, ok := left.(types.Number); ok {
			if r, ok := right.(types.Number); ok {
				value := l.Value + r.Value
				return types.Number{Value: value}, nil
			}
			return nil, newInterpreterError("Cannot add values of different types", expr.Operator)
		}
		if l, ok := left.(types.String); ok {
			if r, ok := right.(types.String); ok {
				value := l.Value + r.Value
				return types.String{Value: value}, nil
			}
			return nil, newInterpreterError("Cannot add values of different types", expr.Operator)
		}
		return nil, newInterpreterError("Can only add string or numeric types", expr.Operator)
	}
	return nil, nil
}

func numeric(args ...any) (bool, string) {
	for _, a := range args {
		if _, ok := a.(types.Number); !ok {
			return false, reflect.TypeOf(a).Name()
		}
	}
	return true, ""
}

func isEqual(equal token.Token, left, right types.ClavType) (bool, error) {
	switch l := left.(type) {
	case types.Number:
		if r, ok := right.(types.Number); ok {
			return l.Value == r.Value, nil
		}
	case types.String:
		if r, ok := right.(types.String); ok {
			return l.Value == r.Value, nil
		}
	case types.Boolean:
		if r, ok := right.(types.Boolean); ok {
			return l.Value == r.Value, nil
		}
	case types.Nil:
		if _, ok := right.(types.Nil); ok {
			return true, nil
		}
	}
	return false, newInterpreterError("Cannot compare values of different types", equal)
}

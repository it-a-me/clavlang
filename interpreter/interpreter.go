package interpreter

import (
	"fmt"
	"reflect"

	"github.com/it-a-me/clavlang/parser"
	"github.com/it-a-me/clavlang/token"
)

func Interpret(expr parser.Expr) error {
	res, err := evaluate(expr)
	if err != nil {
		return err
	}
	fmt.Printf("=%v\n", res)
	return nil
}

func evaluate(expr parser.Expr) (any, error) {
	switch e := expr.(type) {
	case parser.Literal:
		return evalutateLiteral(e), nil
	case parser.Grouping:
		return evalutateGrouping(e)
	case parser.Unary:
		return evalutateUnary(e)
	case parser.Binary:
		return evalutateBinary(e)
	}
	panic("Unreachable")
}

func evalutateLiteral(expr parser.Literal) any {
	return expr.Value
}

func evalutateGrouping(expr parser.Grouping) (any, error) {
	return evaluate(expr.Expression)
}

func evalutateUnary(expr parser.Unary) (any, error) {
	right, err := evaluate(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Operator.Type {
	case token.Bang:
		return !right.(bool), nil
	case token.Minus:
		return -right.(float64), nil
	}
	panic("Unreachable")
}

func evalutateBinary(expr parser.Binary) (any, error) {
	left, err := evaluate(expr.Left)
	if err != nil {
		return nil, err
	}
	right, err := evaluate(expr.Right)
	if err != nil {
		return nil, err
	}
	switch expr.Operator.Type {
	case token.Minus:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot subtract non-numeric type "+t, expr.Operator)
		}
		return left.(float64) - right.(float64), nil
	case token.Slash:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot divide non-numeric type "+t, expr.Operator)
		}
		return left.(float64) / right.(float64), nil
	case token.Star:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot multiply non-numeric type "+t, expr.Operator)
		}
		return left.(float64) * right.(float64), nil
	case token.EqualEqual:
		return isEqual(expr.Operator, left, right)
	case token.BangEqual:
		eq, err := isEqual(expr.Operator, left, right)
		return !eq, err
	case token.Greater:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		return left.(float64) > right.(float64), nil
	case token.GreaterEqual:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		return left.(float64) >= right.(float64), nil
	case token.Less:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		return left.(float64) < right.(float64), nil
	case token.LessEqual:
		if ok, t := numeric(left, right); !ok {
			return nil, newInterpreterError("Cannot order non-numeric type "+t, expr.Operator)
		}
		return left.(float64) <= right.(float64), nil
	case token.Plus:
		if l, ok := left.(float64); ok {
			if r, ok := right.(float64); ok {
				return l + r, nil
			}
			return nil, newInterpreterError("Cannot add values of different types", expr.Operator)
		}
		if l, ok := left.(string); ok {
			if r, ok := right.(string); ok {
				return l + r, nil
			}
			return nil, newInterpreterError("Cannot add values of different types", expr.Operator)
		}
		return nil, newInterpreterError("Can only add string or numeric types", expr.Operator)
	}
	return nil, nil
}

func numeric(args ...any) (bool, string) {
	for _, a := range args {
		if _, ok := a.(float64); !ok {
			return false, reflect.TypeOf(a).Name()
		}
	}
	return true, ""
}

func isEqual(equal token.Token, left, right any) (bool, error) {
	if left == nil && right == nil {
		return true, nil
	}
	switch l := left.(type) {
	case float64:
		if r, ok := right.(float64); ok {
			return l == r, nil
		}
	case string:
		if r, ok := right.(string); ok {
			return l == r, nil
		}
	case bool:
		if r, ok := right.(bool); ok {
			return l == r, nil
		}
	}
	return false, newInterpreterError("Cannot compare values of different types", equal)
}

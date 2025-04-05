package parser

import (
	"github.com/it-a-me/clavlang/token"
	"github.com/it-a-me/clavlang/types"
)

type Parser struct {
	tokens  []token.Token
	current int
}

func NewParser(tokens []token.Token) Parser {
	return Parser{tokens: tokens}
}

func (p *Parser) Parse() ([]Stmt, []error) {
	stmts := []Stmt{}
	var errs []error
	for !p.isAtEnd() {
		s, err := p.declaration()
		if err != nil {
			errs = append(errs, err)
		} else {
			stmts = append(stmts, s)
		}
	}
	return stmts, errs
}

func (p *Parser) declaration() (Stmt, error) {
	if p.match(token.Var) {
		return p.varDeclaration()
	}

	stmt, err := p.statement()
	if err != nil {
		p.synchronize()
		return nil, err
	}
	return stmt, nil
}

func (p *Parser) varDeclaration() (Stmt, error) {
	name, err := p.consume(token.Identifier, "Expect variable name")
	if err != nil {
		return nil, err
	}
	var initializer Expr
	if p.match(token.Equal) {
		initializer, err = p.expression()
		if err != nil {
			return nil, err
		}
	}
	p.consume(token.Semicolon, "Expect ';' after variable declaration")
	return Var{Name: name, Initializer: initializer}, nil
}

func (p *Parser) statement() (Stmt, error) {
	if p.match(token.Print) {
		return p.printStatement()
	}
	return p.expressionStatement()
}

func (p *Parser) printStatement() (Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	if _, err := p.consume(token.Semicolon, "Expect ';' after value"); err != nil {
		return nil, err
	}
	return Print{Inner: value}, nil
}

func (p *Parser) expressionStatement() (Stmt, error) {
	value, err := p.expression()
	if err != nil {
		return nil, err
	}
	if _, err := p.consume(token.Semicolon, "Expect ';' after value"); err != nil {
		return nil, err
	}
	return Expression{Inner: value}, nil
}

func (p *Parser) expression() (Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, err
	}
	for p.match(token.BangEqual, token.EqualEqual) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, err
		}
		expr = Expr(Binary{Left: expr, Operator: operator, Right: right})
	}
	return expr, nil
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, err
	}
	for p.match(token.Greater, token.GreaterEqual, token.Less, token.LessEqual) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		expr = Expr(Binary{Left: expr, Operator: operator, Right: right})
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.match(token.Minus, token.Plus) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		expr = Expr(Binary{Left: expr, Operator: operator, Right: right})
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.match(token.Slash, token.Star) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		expr = Expr(Binary{Left: expr, Operator: operator, Right: right})
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(token.Bang, token.Minus) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, err
		}
		return Expr(Unary{Operator: operator, Right: right}), nil
	}
	return p.primary()
}

func (p *Parser) primary() (Expr, error) {
	switch {
	case p.match(token.False):
		return Expr(Literal{Value: types.Boolean{Value: false}}), nil
	case p.match(token.True):
		return Expr(Literal{Value: types.Boolean{Value: true}}), nil
	case p.match(token.Nil):
		return Expr(Literal{Value: nil}), nil
	case p.match(token.Number, token.String):
		return Expr(Literal{Value: p.previous().Literal}), nil
	case p.match(token.LeftParen):
		expr, err := p.expression()
		if err != nil {
			return nil, err
		}
		_, err = p.consume(token.RightParen, "Expect ')' after expression")
		if err != nil {
			return nil, err
		}
		return Expr(Grouping{Expression: expr}), nil
	case p.match(token.Identifier):
		return Variable{p.previous()}, nil
	}
	return nil, p.newError("Expected Expression")
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type == token.Semicolon {
			return
		}

		switch p.peek().Type {
		case token.Class:
			fallthrough
		case token.Fun:
			fallthrough
		case token.Var:
			fallthrough
		case token.For:
			fallthrough
		case token.If:
			fallthrough
		case token.While:
			fallthrough
		case token.Print:
			fallthrough
		case token.Return:
			return
		default:
			_ = 0
		}

		p.advance()
	}
}

func (p *Parser) consume(tokenType token.Type, orError string) (token.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}
	return token.Token{}, p.newError(orError)
}

func (p *Parser) match(types ...token.Type) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType token.Type) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == token.EOF
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) newError(context string) error {
	return error(ParseError{
		Token:   p.peek(),
		Context: context,
	})
}

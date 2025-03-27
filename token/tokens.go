package token

import "fmt"

//go:generate stringer -type Type
type Type int

const (
	// Single-character tokens.
	LeftParen Type = iota

	RightParen
	LeftBrace
	RightBrace

	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One or two character tokens.
	Bang
	BangEqual

	Equal
	EqualEqual

	Greater
	GreaterEqual

	Less
	LessEqual

	// Literals.
	Identifier
	String
	Number

	// Keywords.
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or

	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)

type Token struct {
	Type    Type
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(tokenType Type, lexeme string, literal any, line int) Token {
	return Token{
		tokenType,
		lexeme,
		literal,
		line,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.Type.String(), t.Lexeme, t.Literal)
}

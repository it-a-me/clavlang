package scanner

import (
	"fmt"

	"github.com/it-a-me/clavlang/token"
)

type Scanner struct {
	source string
	tokens []token.Token

	start   int
	current int
	line    int

	errors []ScanError
}

func NewScanner(source string) Scanner {
	return Scanner{
		source: source,
		line:   1,
	}
}

func (s *Scanner) Scan() ([]token.Token, []ScanError) {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, token.NewToken(token.EOF, "", nil, s.line))
	return s.tokens, s.errors
}

func (s *Scanner) scanToken() *ScanError {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(token.LeftParen, nil)
	case ')':
		s.addToken(token.RightParen, nil)
	case '{':
		s.addToken(token.LeftBrace, nil)
	case '}':
		s.addToken(token.RightBrace, nil)
	case ',':
		s.addToken(token.Comma, nil)
	case '.':
		s.addToken(token.Dot, nil)
	case '-':
		s.addToken(token.Minus, nil)
	case '+':
		s.addToken(token.Plus, nil)
	case ';':
		s.addToken(token.Semicolon, nil)
	case '*':
		s.addToken(token.Star, nil)
	case '!':
		if s.match('=') {
			s.addToken(token.BangEqual, nil)
		} else {
			s.addToken(token.Bang, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EqualEqual, nil)
		} else {
			s.addToken(token.Equal, nil)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LessEqual, nil)
		} else {
			s.addToken(token.Less, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GreaterEqual, nil)
		} else {
			s.addToken(token.Greater, nil)
		}
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.Slash, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.handleString()
	default:
		switch {
		case isDigit(c):
			s.handleNumber()
		case isAlpha(c):
			s.handleIdentifier()
		default:
			return s.NewError(fmt.Sprintf("Unexpected character '%c'", c))
		}
	}
	return nil
}

func (s *Scanner) addToken(tokenType token.TokenType, literal any) {
	lexeme := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token.NewToken(tokenType, lexeme, literal, s.line))
}

func (s *Scanner) advance() byte {
	s.current += 1
	return s.source[s.current-1]
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) handleString() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.NewError("Unterminated string")
		return
	}

	// The closing quote
	s.advance()

	content := s.source[s.start+1 : s.current-1]
	s.addToken(token.String, content)
}

func (s *Scanner) handleNumber() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		for isDigit(s.peek()) {
			s.advance()
		}
	}
}

func (s *Scanner) handleIdentifier() {
	for isAlpha(s.peek()) || isDigit(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	if keyword, exists := keywords[text]; exists {
		s.addToken(keyword, nil)
	} else {
		s.addToken(token.Identifier, nil)
	}
}

func (s *Scanner) NewError(err string) *ScanError {
	e := ScanError{
		s.line,
		err,
	}
	s.errors = append(s.errors, e)
	return &e
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

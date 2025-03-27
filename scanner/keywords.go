package scanner

import "github.com/it-a-me/clavlang/token"

func Keywords(identifier string) (token.Type, bool) {
	keywords := map[string]token.Type{
		"and":    token.And,
		"class":  token.Class,
		"else":   token.Else,
		"false":  token.False,
		"for":    token.For,
		"fun":    token.Fun,
		"if":     token.If,
		"nil":    token.Nil,
		"or":     token.Or,
		"print":  token.Print,
		"return": token.Return,
		"super":  token.Super,
		"this":   token.This,
		"true":   token.True,
		"var":    token.Var,
		"while":  token.While,
	}
	kw, ok := keywords[identifier]
	return kw, ok
}

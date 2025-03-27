package parser

import (
	"fmt"

	"github.com/it-a-me/clavlang/token"
)

type ParseError struct {
	Token   token.Token
	Context string
}

func (p ParseError) Error() string {
	return fmt.Sprintf("Parse Error: %s on line %d.", p.Token.Type.String(), p.Token.Line)
}

package main

import (
	"log"
	"os"
)

const PACKAGE_NAME = "ast"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: generate_ast <output file>")
	}
	content := defineAst(exprTypes)

	err := os.WriteFile(os.Args[1], []byte(content), 0o600)
	if err != nil {
		log.Fatal(err)
	}
}

var exprTypes = []ExprType{
	{"Binary", []string{"left Expr", "operator token.Token", "right Expr"}},
	{"Grouping", []string{"expression Expr"}},
	{"Literal", []string{"value any"}},
	{"Unary", []string{"operator Token", "right Expr"}},
}

type ExprType struct {
	name   string
	fields []string
}

func defineAst(types []ExprType) string {
	content := ""
	content += "package " + PACKAGE_NAME + "\n"
	content += "import \"github.com/it-a-me/clavlang/token\"\n"

	for _, t := range types {
		content += defineType(t)
	}
	return content
}

func defineType(exprType ExprType) string {
	content := "type " + exprType.name + " struct {" + "\n"
	for _, f := range exprType.fields {
		content += "\t" + f + "\n"
	}
	content += "}" + "\n"
	return content
}

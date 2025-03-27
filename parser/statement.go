package parser

type Stmt interface {
	stmt()
}

type Expression struct {
	Inner Expr
}

type Print struct {
	Inner Expr
}

func (Expression) stmt() {}
func (Print) stmt()      {}

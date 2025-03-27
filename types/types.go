package types

type ClavType interface {
	clav()
}

type Number struct {
	Value float64
}

type String struct {
	Value string
}
type Boolean struct {
	Value bool
}

type Nil struct{}

func (Number) clav()  {}
func (String) clav()  {}
func (Boolean) clav() {}
func (Nil) clav()     {}

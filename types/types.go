package types

import "fmt"

type ClavType interface {
	clav()
	String() string
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

func (Number) clav() {}
func (n Number) String() string {
	return fmt.Sprint(n.Value)
}

func (String) clav() {}
func (s String) String() string {
	return s.Value
}
func (Boolean) clav() {}
func (b Boolean) String() string {
	return fmt.Sprint(b.Value)
}
func (Nil) clav() {}
func (n Nil) String() string {
	return "nil"
}

package scanner

import "fmt"

type ScanError struct {
	line int
	err  string
}

func (e ScanError) Error() string {
	return fmt.Sprintf("Error on line %d: %s", e.line, e.err)
}

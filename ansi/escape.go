package ansi

import "fmt"

type EscapeType uint8

const (
	EscapeNone EscapeType = iota
	EscapeZSH
)

func Escape(str string, t EscapeType) string {
	switch t {
	case EscapeNone:
		return str
	case EscapeZSH:
		return fmt.Sprintf("%%{%s%%}", str)
	default:
		panic("bad escape type")
	}
}

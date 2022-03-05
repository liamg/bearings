package ansi

import (
	"fmt"
	"strings"
)

type EscapeType uint8

const (
	EscapeNone EscapeType = iota
	EscapeZSH
)

func EscapeCode(str string, t EscapeType) string {
	switch t {
	case EscapeNone:
		return str
	case EscapeZSH:
		return fmt.Sprintf("%%{%s%%}", str)
	default:
		panic("bad escape type")
	}
}

func EscapeString(str string, t EscapeType) string {
	switch t {
	case EscapeNone:
		return str
	case EscapeZSH:
		return strings.ReplaceAll(str, "%", "%%")
	default:
		panic("bad escape type")
	}
}

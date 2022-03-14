package ansi

import (
	"fmt"
	"strings"

	"github.com/liamg/bearings/state"
)

func EscapeCode(str string, t state.Shell) string {
	switch t {
	case state.ShellZSH:
		return fmt.Sprintf("%%{%s%%}", str)
	case state.ShellBash:
		return fmt.Sprintf("\\[%s\\]", str)
	default:
		return str
	}
}

func EscapeString(str string, t state.Shell) string {
	switch t {
	case state.ShellBash:
		return strings.ReplaceAll(strings.ReplaceAll(str, `\[`, `\\[`), `\]`, `\\]`)
	case state.ShellZSH:
		return strings.ReplaceAll(str, "%", "%%")
	default:
		return str
	}
}

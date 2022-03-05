package powerline

import (
	"fmt"
	"io"

	"github.com/liamg/bearings/ansi"
)

type Writer struct {
	inner  io.Writer
	escape ansi.EscapeType
}

// Powerline runes are at 0xe0b0 -> 0xe0d4
var powerlineRuneMap = make(map[rune]struct{})
var powerlineRunes = []rune{
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
}

/*
var powerlineSeparators = []rune{
	'',
	'',
	'',
	'',
	'',
	'',
	'',
	'',
}
*/

func init() {
	for _, r := range powerlineRunes {
		powerlineRuneMap[r] = struct{}{}
	}
}

func NewWriter(w io.Writer, escape ansi.EscapeType) *Writer {
	return &Writer{
		inner:  w,
		escape: escape,
	}
}

func (p *Writer) write(s string) {
	_, _ = p.inner.Write([]byte(s))
}

func (p *Writer) Reset(str string) {
	p.WriteAnsi("\x1b[0m")
	p.write(str)
}

func (p *Writer) WriteAnsi(str string) {
	p.write(ansi.EscapeCode(str, p.escape))
}

func (p *Writer) Printf(style ansi.Style, shouldEscape bool, format string, args ...interface{}) {
	input := fmt.Sprintf(format, args...)
	if shouldEscape {
		input = ansi.EscapeString(input, p.escape)
	}
	p.write(style.Ansi(p.escape))
	var inverted bool
	for _, r := range input {
		if p.isPowerlineRune(r) != inverted {
			inverted = !inverted
			if inverted {
				p.write(style.SmartInvert().Ansi(p.escape))
			} else {
				p.write(style.Ansi(p.escape))
			}
		}
		p.write(string(r))
	}
}

func (p *Writer) isPowerlineRune(r rune) bool {
	_, ok := powerlineRuneMap[r]
	return ok
}

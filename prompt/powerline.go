package prompt

import (
	"fmt"
	"io"

	"github.com/liamg/bearings/ansi"
)

type powerlineWriter struct {
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

func NewPowerlineWriter(w io.Writer, escape ansi.EscapeType) *powerlineWriter {
	return &powerlineWriter{
		inner:  w,
		escape: escape,
	}
}

func (p *powerlineWriter) write(s string) {
	_, _ = p.inner.Write([]byte(s))
}

func (p *powerlineWriter) Reset() {
	p.write(ansi.Escape("\x1b[0m", p.escape))
}

func (p *powerlineWriter) Printf(style ansi.Style, format string, args ...interface{}) {
	input := fmt.Sprintf(format, args...)
	p.write(style.Ansi(p.escape))
	var inverted bool
	for _, r := range []rune(input) {
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

func (p *powerlineWriter) isPowerlineRune(r rune) bool {
	_, ok := powerlineRuneMap[r]
	return ok
}

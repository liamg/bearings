package powerline

import (
	"fmt"
	"io"

	"github.com/liamg/bearings/ansi"
	"github.com/liamg/bearings/state"
)

type Writer struct {
	inner      io.Writer
	shell      state.Shell
	firstStyle *ansi.Style
	lastStyle  *ansi.Style
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

func NewWriter(w io.Writer, shell state.Shell) *Writer {
	return &Writer{
		inner: w,
		shell: shell,
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
	p.write(ansi.EscapeCode(str, p.shell))
}

func (p *Writer) FirstStyle() *ansi.Style {
	return p.firstStyle
}

func (p *Writer) LastStyle() *ansi.Style {
	return p.lastStyle
}

func (p *Writer) PrintRaw(s string) {
	p.write(s)
}

func (p *Writer) Printf(style ansi.Style, format string, args ...interface{}) {
	input := fmt.Sprintf(format, args...)
	input = ansi.EscapeString(input, p.shell)
	if input == "" {
		return
	}
	if p.firstStyle == nil {
		p.firstStyle = &style
	}
	p.lastStyle = &style
	p.write(style.Ansi(p.shell))
	var inverted bool
	for _, r := range input {
		if p.isPowerlineRune(r) != inverted {
			inverted = !inverted
			if inverted {
				p.write(style.SmartInvert().Ansi(p.shell))
			} else {
				p.write(style.Ansi(p.shell))
			}
		}
		p.write(string(r))
	}
}

func (p *Writer) isPowerlineRune(r rune) bool {
	_, ok := powerlineRuneMap[r]
	return ok
}

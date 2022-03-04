package ansi

import (
	"encoding/hex"
	"fmt"
	"strings"
)

var ansiColourBases = map[string]int{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,
	"default": 9,
}

type Colour struct {
	index int
	rgb   []uint8
	fg    bool
}

var DefaultFg = Colour{
	index: 9,
	fg:    true,
}

var DefaultBg = Colour{
	index: 9,
	fg:    false,
}

func (c Colour) Fg() Colour {
	c.fg = true
	return c
}

func (c Colour) Bg() Colour {
	c.fg = false
	return c
}

func (c Colour) Ansi(a EscapeType) string {
	switch {
	case c.rgb != nil && c.fg:
		return Escape(fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", 38, c.rgb[0], c.rgb[1], c.rgb[2]), a)
	case c.rgb != nil && !c.fg:
		return Escape(fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", 48, c.rgb[0], c.rgb[1], c.rgb[2]), a)
	case c.fg:
		return Escape(fmt.Sprintf("\x1b[%dm", c.index+30), a)
	default:
		return Escape(fmt.Sprintf("\x1b[%dm", c.index+40), a)
	}
}

type Style struct {
	AllowSmartInvert bool
	Foreground       Colour
	Background       Colour
	Bold             bool
	Faint            bool
	Italic           bool
	Underline        bool
	Blink            bool
}

func (f Style) SmartInvert() Style {
	if !f.AllowSmartInvert {
		return f
	}

	return Style{
		Foreground: f.Background.Fg(),
		Background: DefaultBg,
	}
}

func (f Style) WithoutSmartInvert() Style {
	f.AllowSmartInvert = false
	return f
}

func (f Style) WithSmartInvert() Style {
	f.AllowSmartInvert = true
	return f
}

func (f Style) Ansi(a EscapeType) string {
	ansi := f.Background.Ansi(a) + f.Foreground.Ansi(a)
	if f.Bold {
		ansi += "\x1b[1m"
	}
	if f.Faint {
		ansi += "\x1b[2m"
	}
	if f.Italic {
		ansi += "\x1b[3m"
	}
	if f.Underline {
		ansi += "\x1b[4m"
	}
	if f.Blink {
		ansi += "\x1b[5m"
	}
	return ansi
}

func ParseColourString(colour string) Colour {
	if base, ok := ansiColourBases[colour]; ok {
		return Colour{
			index: base,
		}
	}
	return hexToANSI(colour)
}

func hexToANSI(input string) Colour {
	input = strings.TrimPrefix(input, "#")
	if len(input) != 6 {
		return DefaultFg
	}
	raw, err := hex.DecodeString(input)
	if err != nil {
		return DefaultFg
	}
	return Colour{
		index: 0,
		rgb:   raw,
	}
}

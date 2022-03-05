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

func (c Colour) String() string {
	if len(c.rgb) == 3 {
		return fmt.Sprintf("#%02x%02x%02x", c.rgb[0], c.rgb[1], c.rgb[2])
	}
	for name, index := range ansiColourBases {
		if index == c.index {
			return name
		}
	}
	return ""
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
		return EscapeCode(fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", 38, c.rgb[0], c.rgb[1], c.rgb[2]), a)
	case c.rgb != nil && !c.fg:
		return EscapeCode(fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", 48, c.rgb[0], c.rgb[1], c.rgb[2]), a)
	case c.fg:
		return EscapeCode(fmt.Sprintf("\x1b[%dm", c.index+30), a)
	default:
		return EscapeCode(fmt.Sprintf("\x1b[%dm", c.index+40), a)
	}
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

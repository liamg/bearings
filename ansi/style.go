package ansi

import "github.com/liamg/bearings/state"

type Style struct {
	AllowSmartInvert bool
	Foreground       Colour
	Background       Colour
	Bold             bool
	Faint            bool
	Italic           bool
	Underline        bool
	Blink            bool
	From             *Style
}

func (f Style) SmartInvert() Style {
	if !f.AllowSmartInvert {
		return f
	}

	if f.From != nil {
		return Style{
			Foreground: f.From.Background.Fg(),
			Background: f.Background.Bg(),
			Bold:       f.Bold,
			Faint:      f.Faint,
			Underline:  f.Underline,
			Italic:     f.Italic,
			Blink:      f.Blink,
		}
	}

	return Style{
		Foreground: f.Background.Fg(),
		Background: DefaultBg,
		Bold:       f.Bold,
		Faint:      f.Faint,
		Underline:  f.Underline,
		Italic:     f.Italic,
		Blink:      f.Blink,
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

func (f Style) Ansi(a state.Shell) string {
	ansi := f.Background.Ansi(a) + f.Foreground.Ansi(a)
	if f.Bold {
		ansi += EscapeCode("\x1b[1m", a)
	}
	if f.Faint {
		ansi += EscapeCode("\x1b[2m", a)
	}
	if f.Italic {
		ansi += EscapeCode("\x1b[3m", a)
	}
	if f.Underline {
		ansi += EscapeCode("\x1b[4m", a)
	}
	if f.Blink {
		ansi += EscapeCode("\x1b[5m", a)
	}
	return ansi
}

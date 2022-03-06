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
		}
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

func (f Style) Ansi(a state.Shell) string {
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

package prompt

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/liamg/bearings/ansi"
	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/modules"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

func Do(w io.Writer, lastExit int) error {

	wd, _ := os.Getwd()
	home, _ := os.UserHomeDir()

	conf, err := config.Load()
	if err != nil {
		return err
	}
	s := state.State{
		LastExitCode: lastExit,
		WorkingDir:   wd,
		HomeDir:      home,
	}

	switch filepath.Base(os.Getenv("SHELL")) {
	case "zsh":
		s.AnsiEscapeType = ansi.EscapeZSH
	}

	writer := powerline.NewWriter(w, s.AnsiEscapeType)

	style := ansi.Style{
		Foreground: ansi.ParseColourString(conf.Fg).Fg(),
		Background: ansi.ParseColourString(conf.Bg).Bg(),
	}

	writer.Reset("")
	writer.Printf(style, false, strings.Repeat("\n", conf.LinesAbove))

	var lastSep string
	var lastStyle *ansi.Style
	for _, modConf := range conf.Modules {

		buffer := bytes.NewBuffer([]byte{})
		modWriter := powerline.NewWriter(buffer, s.AnsiEscapeType)

		mod, mergedConfig, err := modules.Load(s, conf, modConf)
		if err != nil {
			return err
		}
		mod.Render(modWriter)
		if buffer.Len() == 0 {
			continue
		}
		modStyle := mergedConfig.Style(conf)
		sepStyle := modStyle
		sepStyle.From = lastStyle
		writer.Printf(sepStyle.WithSmartInvert(), true, "%s", lastSep)
		content := strings.ReplaceAll(mergedConfig.Label(), "%s", buffer.String())
		writer.Printf(modStyle, false, " %s ", content)
		lastSep = mergedConfig.String("divider", conf.Divider)
		lastStyle = &modStyle
	}

	if lastStyle != nil {
		style = *lastStyle
	}
	writer.Printf(style.WithSmartInvert(), true, "%s", conf.End)
	writer.Reset(" ")
	return nil
}

package prompt

import (
	"bytes"
	"fmt"
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

	sepStyle := ansi.Style{
		Foreground: ansi.ParseColourString(conf.DividerFg).Fg(),
		Background: style.Background,
	}

	writer.Reset("")
	writer.Printf(style, strings.Repeat("\n", conf.LinesAbove)+" ")

	var lastSep string
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
		writer.Printf(sepStyle.WithSmartInvert(), "%s", lastSep)
		writer.PrintfWithLabel(modStyle, mergedConfig.Label(), "%s", buffer.String())
		lastSep = fmt.Sprintf(" %s ", mergedConfig.String("divider", conf.Divider))
	}

	writer.Printf(style.WithSmartInvert(), " %s", conf.End)
	writer.Reset(" ")
	return nil
}

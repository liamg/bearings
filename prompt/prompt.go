package prompt

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/liamg/bearings/ansi"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/modules"
	"github.com/liamg/bearings/state"
)

func Do(w io.Writer, lastExit int) error {

	var ansiEscape ansi.EscapeType

	switch filepath.Base(os.Getenv("SHELL")) {
	case "zsh":
		ansiEscape = ansi.EscapeZSH
	}

	writer := NewPowerlineWriter(w, ansiEscape)

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

	style := ansi.Style{
		Foreground: ansi.ParseColourString(conf.Fg).Fg(),
		Background: ansi.ParseColourString(conf.Bg).Bg(),
	}

	sepStyle := ansi.Style{
		Foreground: ansi.ParseColourString(conf.DividerFg).Fg(),
		Background: style.Background,
	}

	writer.Reset()
	writer.Printf(style, strings.Repeat("\n", conf.LinesAbove)+" ")

	var lastSep string
	for _, modConf := range conf.Modules {

		mod, mergedConfig, err := modules.Load(s, conf, modConf)
		if err != nil {
			return err
		}
		content := mod.Render()
		if content == "" {
			continue
		}
		content = strings.Replace(mergedConfig.Label(), "%s", content, 1)
		modStyle := mergedConfig.Style(conf)
		writer.Printf(sepStyle.WithSmartInvert(), "%s", lastSep)
		writer.Printf(modStyle.WithoutSmartInvert(), "%s", content)

		lastSep = fmt.Sprintf(" %s ", mergedConfig.String("separator", conf.Divider))
	}

	writer.Printf(style.WithSmartInvert(), " %s", conf.End)
	writer.Reset()
	writer.write(" ")
	return nil
}

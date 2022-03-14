package prompt

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/liamg/bearings/ansi"
	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/modules"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

func Do(w io.Writer, lastExit int, forceShell string, lastDuration float64, jobCount int) error {

	conf, err := config.Load()
	if err != nil {
		return err
	}

	s := state.Derive(lastExit, forceShell, lastDuration, jobCount)

	writer := powerline.NewWriter(w, s.Shell)

	style := ansi.Style{
		Foreground: ansi.ParseColourString(conf.Fg).Fg(),
		Background: ansi.ParseColourString(conf.Bg).Bg(),
	}

	writer.Reset("")
	writer.Printf(style, strings.Repeat("\n", conf.LinesAbove))

	var lastSep string
	var lastStyle *ansi.Style
	for _, modConf := range conf.Modules {

		buffer := bytes.NewBuffer([]byte{})
		modWriter := powerline.NewWriter(buffer, s.Shell)

		mod, mergedConfig, err := modules.Load(s, conf, modConf)
		if err != nil {
			return err
		}
		mod.Render(modWriter)
		if buffer.Len() == 0 {
			continue
		}
		modStyle := mergedConfig.Style(conf)
		modStyle.From = lastStyle
		writer.Printf(modStyle.WithSmartInvert(), "%s", lastSep)
		if first := modWriter.FirstStyle(); first != nil {
			writer.PrintRaw(first.Ansi(s.Shell))
		}
		content := strings.ReplaceAll(mergedConfig.Label(), "%s", buffer.String())
		lastSep = mergedConfig.String("divider", conf.Divider)
		lastStyle = modWriter.LastStyle()
		paddingBefore := strings.Repeat(" ", mergedConfig.Int("padding_before", conf.Padding))
		paddingAfter := strings.Repeat(" ", mergedConfig.Int("padding_after", conf.Padding))
		writer.PrintRaw(fmt.Sprintf("%s%s%s", paddingBefore, content, paddingAfter))
	}
	if lastStyle != nil {
		style = *lastStyle
	}
	writer.Printf(style.WithSmartInvert(), "%s", conf.End)
	writer.Reset(" ")
	writer.WriteAnsi("\x1b[0K\x1b[0m")
	return nil
}

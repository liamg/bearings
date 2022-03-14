package modules

import (
	"github.com/liamg/bearings/ansi"
	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type exitCodeModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

const (
	iconExitSuccess = ""
	iconExitFailure = ""
)

func init() {
	register("exitcode", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &exitCodeModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label":          "%s",
		"show_success":   false,
		"show_exit_code": true,
		"success_bg":     "",
		"failure_bg":     "",
		"success_fg":     "green",
		"failure_fg":     "red",
		"success_output": iconExitSuccess,
		"failure_output": iconExitFailure,
	})
}

func (e *exitCodeModule) Render(w *powerline.Writer) {
	baseStyle := e.mc.Style(e.gc)
	if e.state.LastExitCode > 0 {
		baseStyle.Foreground = ansi.ParseColourString(
			e.mc.String(
				"failure_fg",
				"red",
			),
		).Fg()
		bg := e.mc.String("failure_bg", "")
		if bg != "" {
			baseStyle.Background = ansi.ParseColourString(
				e.mc.String(
					"failure_bg",
					e.gc.Bg,
				),
			).Bg()
		}
		if e.mc.Bool("show_exit_code", true) {
			w.Printf(
				baseStyle,
				"%s %d",
				e.mc.String("failure_output", iconExitFailure),
				e.state.LastExitCode,
			)
		} else {
			w.Printf(
				baseStyle,
				"%s",
				e.mc.String("failure_output", iconExitFailure),
			)
		}
	} else if e.mc.Bool("show_success", false) {
		baseStyle.Foreground = ansi.ParseColourString(
			e.mc.String(
				"success_fg",
				"green",
			),
		).Fg()
		bg := e.mc.String("success_bg", "")
		if bg != "" {
			baseStyle.Background = ansi.ParseColourString(
				e.mc.String(
					"success_bg",
					e.gc.Bg,
				),
			).Bg()
		}
		w.Printf(
			baseStyle,
			"%s",
			e.mc.String("success_output", iconExitSuccess),
		)
	}
}

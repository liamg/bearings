package modules

import (
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
		"success_output": iconExitSuccess,
		"failure_output": iconExitFailure,
	})
}

func (e *exitCodeModule) Render(w *powerline.Writer) {
	baseStyle := e.mc.Style(e.gc)
	if e.state.LastExitCode > 0 {
		w.Printf(
			baseStyle,
			"%s %d",
			e.mc.String("failure_output", iconExitFailure),
			e.state.LastExitCode,
		)
	} else if e.mc.Bool("show_sucess", false) {
		w.Printf(
			baseStyle,
			"%s",
			e.mc.String("success_output", iconExitSuccess),
		)
	}
}

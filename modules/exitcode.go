package modules

import (
	"fmt"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/state"
)

type exitCodeModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("exitcode", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &exitCodeModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"fg": "red",
	})
}

func (e *exitCodeModule) Render() string {
	if e.state.LastExitCode > 0 {
		return fmt.Sprintf(
			"\uF071 %d",
			e.state.LastExitCode)
	}
	return ""
}

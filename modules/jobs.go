package modules

import (
	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type jobsModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("jobs", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &jobsModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label": "%s jobs",
	})
}

func (e *jobsModule) Render(w *powerline.Writer) {
	if e.state.JobCount == 0 {
		return
	}
	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "%d", e.state.JobCount)
}

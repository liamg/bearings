package modules

import (
	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type newlineModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("newline", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &newlineModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label":          "%s",
		"padding_before": 0,
		"padding_after":  0,
	})
}

func (e *newlineModule) Render(w *powerline.Writer) {
	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "\n")
}

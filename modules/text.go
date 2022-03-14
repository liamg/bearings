package modules

import (
	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type textModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("text", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &textModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label": "%s",
		"text":  "",
	})
}

func (e *textModule) Render(w *powerline.Writer) bool {
	text := e.mc.String("text", "")
	if text == "" {
		return false
	}
	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "%s", text)
	return false
}

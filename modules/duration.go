package modules

import (
	"time"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type durationModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("duration", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &durationModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label": "%s",
	})
}

func (e *durationModule) Render(w *powerline.Writer) {
	var duration time.Duration
	switch e.state.Shell {
	case state.ShellFish, state.ShellBash:
		duration = time.Millisecond * time.Duration(e.state.LastDuration)
	case state.ShellZSH:
		duration = time.Millisecond * time.Duration(1000*e.state.LastDuration)
	default:
		duration = time.Second * time.Duration(e.state.LastDuration)
	}
	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "%s", duration)
}

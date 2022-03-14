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
		"label":     "%s",
		"threshold": "3s",
	})
}

func (e *durationModule) Render(w *powerline.Writer) {
	var duration time.Duration
	switch e.state.Shell {
	case state.ShellFish, state.ShellBash, state.ShellZSH:
		duration = time.Millisecond * time.Duration(e.state.LastDuration)
	default:
		duration = time.Second * time.Duration(e.state.LastDuration)
	}
	threshold := e.mc.String("threshold", "-1s")
	t, err := time.ParseDuration(threshold)
	if err != nil {
		t = time.Second * -1
	}
	if duration < t {
		return
	}
	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "%s", duration)
}

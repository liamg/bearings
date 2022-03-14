package modules

import (
	"os"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type hostnameModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("hostname", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &hostnameModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label": "%s",
	})
}

func (e *hostnameModule) Render(w *powerline.Writer) bool {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "???"
	}
	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "%s", hostname)
	return false
}

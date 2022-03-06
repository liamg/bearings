package modules

import (
	"os/user"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type usernameModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("username", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &usernameModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label": "%s",
	})
}

func (e *usernameModule) Render(w *powerline.Writer) {
	username := "?"
	if u, err := user.Current(); err == nil {
		username = u.Username
	}
	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "%s", username)
}

package modules

import (
	"path/filepath"
	"strings"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/state"
)

type workDirModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("workdir", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &workDirModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label": "\uE613 %s",
	})
}

func (e *workDirModule) Render() string {
	clean := strings.TrimPrefix(e.state.WorkingDir, e.state.HomeDir)
	if clean != e.state.WorkingDir {
		clean = filepath.Join("~", clean)
	}
	return clean
}

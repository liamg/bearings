package modules

import (
	"path/filepath"
	"strings"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type workDirModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("cwd", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &workDirModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label":     "%s",
		"max_depth": 0,
	})
}

func (e *workDirModule) Render(w *powerline.Writer) {
	baseStyle := e.mc.Style(e.gc)
	clean := strings.TrimPrefix(e.state.WorkingDir, e.state.HomeDir)
	if clean != e.state.WorkingDir {
		clean = filepath.Join("~", clean)
	}
	if max := e.mc.Int("max_depth", 0); max > 0 {
		parts := strings.Split(clean, string(filepath.Separator))
		if len(parts) > max {
			clean = filepath.Join(append([]string{"..."}, parts[len(parts)-max:]...)...)
		}
	}
	w.Printf(
		baseStyle,
		"%s",
		clean,
	)
}
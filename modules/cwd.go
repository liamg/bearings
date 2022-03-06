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
		"label":        "%s",
		"max_depth":    0,
		"separator":    " \uE0B1 ",
		"separator_fg": "",
	})
}

func (e *workDirModule) Render(w *powerline.Writer) {
	clean := strings.TrimPrefix(e.state.WorkingDir, e.state.HomeDir)
	if clean != e.state.WorkingDir {
		clean = filepath.Join("~", clean)
	}
	parts := strings.Split(clean, string(filepath.Separator))
	if max := e.mc.Int("max_depth", 0); max > 0 {
		if len(parts) > max {
			parts = append([]string{"..."}, parts[len(parts)-max:]...)
		}
	}
	baseStyle := e.mc.Style(e.gc)
	sepStyle := baseStyle
	fg := e.mc.String("separator_fg", "")
	if fg != "" {
		sepStyle.Foreground = e.mc.Fg("separator_fg", baseStyle.Foreground.String())
	}
	separator := e.mc.String("separator", string(filepath.Separator))
	for i, part := range parts {
		if i > 0 {
			w.Printf(sepStyle, "%s", separator)
		}
		w.Printf(baseStyle, "%s", part)
	}
}

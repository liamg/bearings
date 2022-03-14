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
		"label":       "%s",
		"home_text":   "~",
		"max_depth":   0,
		"separator":   " \uE0B1 ",
		"deep_prefix": "\uF141",
	})
}

func (e *workDirModule) Render(w *powerline.Writer) bool {
	clean := strings.TrimPrefix(e.state.WorkingDir, e.state.HomeDir)
	if clean != e.state.WorkingDir {
		clean = filepath.Join(e.mc.String("home_text", "~"), clean)
	}
	parts := strings.Split(clean, string(filepath.Separator))
	if max := e.mc.Int("max_depth", 0); max > 0 {
		if len(parts) > max {
			parts = append([]string{e.mc.String("deep_prefix", "...")}, parts[len(parts)-(max):]...)
		}
	}
	baseStyle := e.mc.Style(e.gc)
	sepStyle := baseStyle
	sepStyle.Foreground = e.mc.Fg("separator_fg", baseStyle.Foreground.String())
	separator := e.mc.String("separator", string(filepath.Separator))
	for i, part := range parts {
		if i > 0 || e.mc.Bool("separator_at_start", false) {
			w.Printf(sepStyle, "%s", separator)
		}
		w.Printf(baseStyle, "%s", part)
	}
	return false
}

package modules

import (
	"os"
	"sort"
	"strings"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type languagesModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

var languageIcons = map[string]string{
	"go.mod":       "ﳑ",
	"Dockerfile":   "",
	"package.json": "",
	"build.gradle": "",
	"pom.xml":      "",
	"Gemfile":      "",
}

func init() {
	register("languages", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &languagesModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label": "%s",
	})
}

func (e *languagesModule) Render(w *powerline.Writer) {
	baseStyle := e.mc.Style(e.gc)
	var icons []string
	for filename, icon := range languageIcons {
		if _, err := os.Stat(filename); err == nil {
			icons = append(icons, icon)
		}
	}
	sort.Strings(icons)
	w.Printf(baseStyle, "%s", strings.Join(icons, " "))
}

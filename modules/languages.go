package modules

import (
	"fmt"
	"os"
	"sort"

	"github.com/liamg/bearings/ansi"
	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type languagesModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

type icon struct {
	config string
	glyph  string
	colour string
}

var languageIcons = map[string]icon{
	"go.mod":       {"go", "ﳑ", "lightblue"},
	"Dockerfile":   {"docker", "", "blue"},
	"package.json": {"node", "", "green"},
	"build.gradle": {"gradle", "", "green"},
	"pom.xml":      {"pom", "", "blue"},
	"Gemfile":      {"ruby", "", "red"},
}

func init() {
	register("languages", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &languagesModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label":     "%s",
		"separator": " ",
	})
}

func (e *languagesModule) Render(w *powerline.Writer) bool {
	baseStyle := e.mc.Style(e.gc)
	var icons []icon
	for filename, icon := range languageIcons {
		if _, err := os.Stat(filename); err == nil {
			icons = append(icons, icon)
		}
	}

	sort.Slice(icons, func(i, j int) bool {
		return icons[i].glyph < icons[j].glyph
	})

	for i, icon := range icons {
		iconStyle := baseStyle
		iconStyle.Foreground = ansi.ParseColourString(e.mc.String(fmt.Sprintf("colour_%s", icon.config), icon.colour)).Fg()
		separator := e.mc.String("separator", " ")
		if i == len(icons)-1 {
			separator = ""
		}
		w.Printf(iconStyle, "%s%s", icon.glyph, separator)
	}
	return false
}

package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/liamg/bearings/ansi"
)

type ModuleConfig map[string]interface{}

func (c ModuleConfig) Merge(overrides ModuleConfig) ModuleConfig {
	for key, val := range overrides {
		c[key] = val
	}
	return c
}

func (c ModuleConfig) Type() string {
	return c.String("type", "")
}

func (c ModuleConfig) Label() string {
	return c.String("label", "%s")
}

func (c ModuleConfig) Fg(name string, def string) ansi.Colour {
	colour := def
	if val, ok := c[name]; ok {
		colour = fmt.Sprintf("%s", val)
	}
	return ansi.ParseColourString(colour).Fg()
}

func (c ModuleConfig) Bg(name string, def string) ansi.Colour {
	colour := def
	if val, ok := c[name]; ok {
		colour = fmt.Sprintf("%s", val)
	}
	return ansi.ParseColourString(colour).Bg()
}

func (c ModuleConfig) Style(inherit *Config) ansi.Style {
	return ansi.Style{
		AllowSmartInvert: true,
		Foreground:       c.Fg("fg", inherit.Fg),
		Background:       c.Bg("bg", inherit.Bg),
	}
}

func (c ModuleConfig) String(name string, orDefault string) string {
	if val, ok := c[name]; ok {
		return fmt.Sprintf("%s", val)
	}
	return orDefault
}

func (c ModuleConfig) Bool(name string, orDefault bool) bool {
	if val, ok := c[name]; ok {
		switch v := val.(type) {
		case string:
			return strings.EqualFold(v, "true")
		case bool:
			return v
		case int:
			return v > 0
		case float32:
			return v > 0
		case float64:
			return v > 0
		}
	}
	return orDefault
}

func (c ModuleConfig) Int(name string, orDefault int) int {
	if val, ok := c[name]; ok {
		switch v := val.(type) {
		case string:
			if i, err := strconv.Atoi(v); err == nil {
				return i
			}
		case float32:
			return (int)(v)
		case float64:
			return (int)(v)
		case int:
			return v
		}
	}
	return orDefault
}

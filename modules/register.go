package modules

import (
	"fmt"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/state"
)

type registration struct {
	factory  ModuleFactory
	defaults config.ModuleConfig
}

type ModuleFactory func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error)

var registry = make(map[string]registration)

func register(name string, factory ModuleFactory, defaults config.ModuleConfig) {
	registry[name] = registration{
		factory:  factory,
		defaults: defaults,
	}
}

func Load(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, config.ModuleConfig, error) {
	name := mc.Type()
	if reg, ok := registry[name]; ok {
		merged := reg.defaults.Merge(mc)
		built, err := reg.factory(state, gc, merged)
		if err != nil {
			return nil, nil, err
		}
		return built, merged, nil
	}
	return nil, nil, fmt.Errorf("module '%s' not found", name)
}

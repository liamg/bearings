package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func Load() (*Config, error) {
	confPath, err := path()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(confPath)
	if err != nil {
		if os.IsNotExist(err) {
			// write a default config to file for the user to edit in future
			if err := os.MkdirAll(filepath.Dir(confPath), 0700); err == nil {
				if w, err := os.Create(confPath); err == nil {
					_ = yaml.NewEncoder(w).Encode(DefaultConfig)
					_ = w.Close()
				}
			}
			return &DefaultConfig, nil
		}
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var conf Config
	if err := yaml.NewDecoder(f).Decode(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

package config

var DefaultConfig = Config{
	Bg:         "black",
	Fg:         "white",
	End:        "\ue0b0",
	Divider:    "\ue0b0",
	LinesAbove: 1,
	Padding:    1,
	Modules: []ModuleConfig{
		{"type": "exitcode"},
		{"type": "cwd"},
		{"type": "git"},
	},
}

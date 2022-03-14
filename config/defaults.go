package config

var DefaultConfig = Config{
	Bg:         "black",
	Fg:         "white",
	End:        "\ue0b0",
	Divider:    "\ue0b0",
	LinesAbove: 1,
	Padding:    1,
	Modules: []ModuleConfig{
		{
			"type":           "exitcode",
			"show_success":   true,
			"success_output": "",
			"success_fg":     "#ffffff",
			"success_bg":     "#000000",
			"failure_fg":     "#ffffff",
			"failure_bg":     "#bb4444",
		},
		{
			"type":      "duration",
			"threshold": "3s",
			"fg":        "#334488",
			"bg":        "#ffffff",
		},
		{
			"type":      "cwd",
			"label":     " %s",
			"max_depth": 3,
			"fg":        "#aaaaaa",
			"bg":        "#334488",
		},
		{
			"type": "git",
			"fg":   "#777777",
			"bg":   "#393939",
		},
	},
}

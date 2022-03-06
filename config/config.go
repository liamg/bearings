package config

type Config struct {
	Padding    int            `yaml:"padding"`
	End        string         `yaml:"end"`
	Divider    string         `yaml:"divider"`
	DividerFg  string         `yaml:"divider_fg"`
	Fg         string         `yaml:"fg"`
	Bg         string         `yaml:"bg"`
	LinesAbove int            `yaml:"lines_above"`
	Modules    []ModuleConfig `yaml:"modules"`
}

package state

import (
	"os"
	"path/filepath"

	"github.com/liamg/bearings/ansi"
)

type State struct {
	AnsiEscapeType ansi.EscapeType
	LastExitCode   int
	WorkingDir     string
	HomeDir        string
	ShellPath      string
	Shell          string
}

func Derive(lastExit int) State {
	wd, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	s := State{
		LastExitCode: lastExit,
		WorkingDir:   wd,
		HomeDir:      home,
		ShellPath:    os.Getenv("SHELL"),
	}
	s.Shell = filepath.Base(s.ShellPath)
	switch s.Shell {
	case "zsh":
		s.AnsiEscapeType = ansi.EscapeZSH
	}
	return s
}

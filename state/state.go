package state

import (
	"os"
	"path/filepath"
)

type Shell string

const (
	ShellZSH  = "zsh"
	ShellBash = "bash"
	ShellFish = "fish"
)

type State struct {
	LastExitCode int
	WorkingDir   string
	HomeDir      string
	ShellPath    string
	Shell        Shell
}

func Derive(lastExit int, forceShell string) State {
	wd, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	s := State{
		LastExitCode: lastExit,
		WorkingDir:   wd,
		HomeDir:      home,
		ShellPath:    os.Getenv("SHELL"),
		Shell:        Shell(forceShell),
	}
	if forceShell == "" {
		s.Shell = Shell(filepath.Base(s.ShellPath))
	}
	return s
}

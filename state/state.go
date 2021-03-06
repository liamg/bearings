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
	LastDuration float64
	WorkingDir   string
	HomeDir      string
	ShellPath    string
	Shell        Shell
	JobCount     int
}

func Derive(lastExit int, forceShell string, lastDuration float64, jobCount int) State {
	wd, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	s := State{
		LastExitCode: lastExit,
		LastDuration: lastDuration,
		WorkingDir:   wd,
		HomeDir:      home,
		ShellPath:    os.Getenv("SHELL"),
		Shell:        Shell(forceShell),
		JobCount:     jobCount,
	}
	if forceShell == "" {
		s.Shell = Shell(filepath.Base(s.ShellPath))
	}
	return s
}

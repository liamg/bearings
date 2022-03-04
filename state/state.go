package state

import "github.com/liamg/bearings/ansi"

type State struct {
	AnsiEscapeType ansi.EscapeType
	LastExitCode   int
	WorkingDir     string
	HomeDir        string
}

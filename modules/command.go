package modules

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/liamg/bearings/config"
	"github.com/liamg/bearings/powerline"
	"github.com/liamg/bearings/state"
)

type commandModule struct {
	state state.State
	gc    *config.Config
	mc    config.ModuleConfig
}

func init() {
	register("command", func(state state.State, gc *config.Config, mc config.ModuleConfig) (Module, error) {
		return &commandModule{
			state: state,
			mc:    mc,
			gc:    gc,
		}, nil
	}, config.ModuleConfig{
		"label":   "%s",
		"command": "",
	})
}

func (e *commandModule) Render(w *powerline.Writer) {
	command := e.mc.String("command", "")
	if command == "" {
		return
	}

	buffer := bytes.NewBuffer([]byte{})

	cmd := exec.Command(e.state.ShellPath)
	cmd.Stdin = strings.NewReader(command)
	cmd.Stdout = buffer
	cmd.Stderr = buffer
	_ = cmd.Run()

	baseStyle := e.mc.Style(e.gc)
	w.Printf(baseStyle, "%s", strings.TrimSpace(buffer.String()))
}

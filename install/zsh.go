package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	startMarker = "#bearings-auto:start"
	endMarker   = "#bearings-auto:end"
)

func installZSH() error {

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot determine home directory: %w", err)
	}

	target := filepath.Join(home, ".zshrc")

	var entireConfig string
	if data, err := os.ReadFile(target); err == nil {
		entireConfig = string(data)
	}

	before := strings.Split(entireConfig, startMarker)[0]
	after := "\n"
	if strings.Contains(entireConfig, endMarker) {
		after = strings.Split(entireConfig, endMarker)[1]
	}

	injection := `
function configure_bearings() {
    PROMPT="$(bearings prompt -e $?)"
}
[ ! "$TERM" = "linux" ] && precmd_functions+=(configure_bearings)
`

	/*
		prompt() { export PROMPT=$(bearings prompt -e $?); }
		export PROMPT_COMMAND=prompt
		precmd() { prompt; }
		#bearings-auto:end
	*/

	rebuilt := fmt.Sprintf("%s%s%s%s%s", before, startMarker, injection, endMarker, after)
	return os.WriteFile(target, []byte(rebuilt), 0600)
}

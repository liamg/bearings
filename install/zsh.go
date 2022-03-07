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
zmodload zsh/datetime
BEARINGS_TIMER=$EPOCHREALTIME
function preexec() {
  BEARINGS_TIMER=$EPOCHREALTIME
}
function configure_bearings() {
    local DURATION="$(($EPOCHREALTIME - $BEARINGS_TIMER))"
    PROMPT="$(bearings prompt -s zsh -e $? -d $DURATION)"
}
[ ! "$TERM" = "linux" ] && precmd_functions+=(configure_bearings)
`

	rebuilt := fmt.Sprintf("%s%s%s%s%s", before, startMarker, injection, endMarker, after)
	return os.WriteFile(target, []byte(rebuilt), 0600)
}

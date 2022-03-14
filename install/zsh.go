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
function preexec() {
  btimer=$(($(date +%s%0N)/1000000))
}
function configure_bearings() {
    elapsed=0
    if [ $btimer ]; then
      now=$(($(date +%s%0N)/1000000))
      elapsed=$(($now-$btimer))
      unset btimer
    fi
    PROMPT="$(bearings prompt -s zsh -e $? -d ${elapsed} -j $(jobs | wc -l))"
}
[ ! "$TERM" = "linux" ] && precmd_functions+=(configure_bearings)
`

	rebuilt := fmt.Sprintf("%s%s%s%s%s", before, startMarker, injection, endMarker, after)
	return os.WriteFile(target, []byte(rebuilt), 0600)
}

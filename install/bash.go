package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func installBash() error {

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot determine home directory: %w", err)
	}

	target := filepath.Join(home, ".bashrc")

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
PS0='$(echo "$(($(date +%s%N)/1000000))" > /tmp/bearings.$$)';
bearings_prompt() { 
    NOW=$(($(date +%s%N)/1000000))
    START=$NOW
    [[ -f /tmp/bearings.$$ ]] && START=$(cat /tmp/bearings.$$) && rm /tmp/bearings.$$
    DURATION=$(($NOW - $START));
    export PS1=$(bearings prompt -s bash -e $? -d $DURATION); 
}
[[ ! "$TERM" = "linux" ]] && export PROMPT_COMMAND=bearings_prompt
`

	rebuilt := fmt.Sprintf("%s%s%s%s%s", before, startMarker, injection, endMarker, after)
	return os.WriteFile(target, []byte(rebuilt), 0600)
}

package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/liamg/bearings/config"
)

func installFish() error {

	path, err := config.Root()
	if err != nil {
		return fmt.Errorf("cannot determine config directory: %w", err)
	}

	target := filepath.Join(path, "fish", "config.fish")

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
function fish_prompt
    bearings prompt -s fish -e $status -d $CMD_DURATION
end
`

	rebuilt := fmt.Sprintf("%s%s%s%s%s", before, startMarker, injection, endMarker, after)
	if err := os.MkdirAll(filepath.Dir(target), 0700); err != nil {
		return err
	}
	return os.WriteFile(target, []byte(rebuilt), 0600)
}

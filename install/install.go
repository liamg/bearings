package install

import (
	"fmt"
	"log"

	"github.com/liamg/bearings/state"
)

func Do(shell state.Shell) error {

	log.Printf("Shell detected as '%s'.", shell)

	switch shell {
	case state.ShellZSH:
		if err := installZSH(); err != nil {
			return err
		}
	case state.ShellBash:
		if err := installBash(); err != nil {
			return err
		}
	case state.ShellFish:
		if err := installFish(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("shell '%s' is not supported - please configure your shell manually", shell)
	}

	log.Printf("Configuration successful! Please start a new shell in order for it to take effect.")
	return nil
}

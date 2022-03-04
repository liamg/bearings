package install

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Do() error {

	shell := filepath.Base(os.Getenv("SHELL"))

	log.Printf("Shell detected as '%s'.", shell)

	switch shell {
	case "zsh":
		if err := installZSH(); err != nil {
			return err
		}
	case "bash":
		if err := installBash(); err != nil {
			return err
		}
	case "fish":
		if err := installFish(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("shell '%s' is not supported - please configure your shell manually", shell)
	}

	log.Printf("Configuration successful! Please start a new shell in order for it to take effect.")
	return nil
}

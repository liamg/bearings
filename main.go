package main

import (
	"fmt"
	"os"

	"github.com/liamg/bearings/prompt"

	"github.com/spf13/cobra"
)

var flagLastExitCode int

var rootCmd = &cobra.Command{
	Use:           "bearings",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return prompt.Do(cmd.OutOrStdout(), flagLastExitCode)
	},
}

func main() {
	rootCmd.Flags().IntVarP(&flagLastExitCode, "exit", "e", flagLastExitCode, "Last exit code. Should be supplied via $?.")
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/liamg/bearings/install"
	"github.com/liamg/bearings/prompt"
	"github.com/liamg/bearings/state"

	"github.com/spf13/cobra"
)

var flagLastExitCode int
var flagShell string

var rootCmd = &cobra.Command{
	Use: "bearings",
}

var promptCmd = &cobra.Command{
	Use:           "prompt",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmd.SilenceUsage = true
		return prompt.Do(cmd.OutOrStdout(), flagLastExitCode, flagShell)
	},
}

var installCmd = &cobra.Command{
	Use:           "install",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmd.SilenceUsage = true
		s := state.Derive(0, flagShell)
		return install.Do(s.Shell)
	},
}

func main() {
	promptCmd.Flags().
		IntVarP(&flagLastExitCode, "exit", "e", flagLastExitCode, "Last exit code. Should be supplied via $?.")
	rootCmd.PersistentFlags().
		StringVarP(&flagShell, "shell", "s", flagShell, "Shell to install bearings for. Auto-detects by default.")
	rootCmd.AddCommand(promptCmd)
	rootCmd.AddCommand(installCmd)
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

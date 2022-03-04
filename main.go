package main

import (
	"fmt"
	"os"

	"github.com/liamg/bearings/install"
	"github.com/liamg/bearings/prompt"

	"github.com/spf13/cobra"
)

var flagLastExitCode int

var rootCmd = &cobra.Command{
	Use: "bearings",
}

var promptCmd = &cobra.Command{
	Use:           "prompt",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmd.SilenceUsage = true
		return prompt.Do(cmd.OutOrStdout(), flagLastExitCode)
	},
}

var installCmd = &cobra.Command{
	Use:           "install",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmd.SilenceUsage = true
		return install.Do()
	},
}

func main() {
	promptCmd.Flags().
		IntVarP(&flagLastExitCode, "exit", "e", flagLastExitCode, "Last exit code. Should be supplied via $?.")
	rootCmd.AddCommand(promptCmd)
	rootCmd.AddCommand(installCmd)
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

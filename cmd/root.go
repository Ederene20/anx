package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "anx",
		Short: "Find a solution to your bugs faster without opening your browser - Powered by ChatGPT",
		Long: `Anx is a simple cli tool that allow you to interact with ChatGPT from your terminal and solve
				your bugs faster.`,
	}
)

// Executes the root command
func Execute() error {
	return rootCmd.Execute()
}

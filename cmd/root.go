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
		Long: `A simple cli tool that will allow you to interact with ChatGPT from your terminal and solve
				your bugs faster`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("Anx - Powered by ChatGPT")
		//},
	}
)

// Executes the root command
func Execute() error {
	return rootCmd.Execute()
}

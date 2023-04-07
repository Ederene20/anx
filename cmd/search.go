package cmd

import (
	"fmt"
	"strings"

	"github.com/Ederene20/anx/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Ask something to ChaptGPT",
	Long:  `This command allow to have a chat with ChatGPT`,
	Run: func(cmd *cobra.Command, args []string) {

		response := api.MakeRequestToChatGPT(strings.Join(args, " "))

		fmt.Println(response)
	},
}

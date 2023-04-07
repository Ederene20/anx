package cmd

import (
	"fmt"
	"strings"

	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setApiKeyCmd)
}

var setApiKeyCmd = &cobra.Command{
	Use:   "set-api-key",
	Short: "Set OpenAI API KEY",
	Long:  `This command allow to set your OpenAI API KEY to be able to use Anx.`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := exec.Command("/bin/bash", "cmd/set-api-key.sh", strings.Join(args, " ")).Output()

		if err != nil {
			fmt.Printf("error %s", err)
			fmt.Println(output)
		}

		fmt.Println("API Key successfully added.")
	},
}

package cmd

import (
	"fmt"
	"os"
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
		set_api_key_script_location := os.Getenv("ANX_INSTALL") + "/bin/set-api-key.sh"
		output, err := exec.Command("/bin/bash", set_api_key_script_location, strings.Join(args, " ")).Output()

		if err != nil {
			fmt.Printf("error %s", err)
			fmt.Println(output)
		}

		fmt.Println("API Key successfully added.")
	},
}

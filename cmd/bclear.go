package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var bclearCmd = &cobra.Command{
	Use:   "bclear",
	Short: "This utility is used to remove the badge in iTerm2 created by the [tabset](https://github.com/jonathaneunice/iterm2-tab-set) utility.",
	Run: func(cmd *cobra.Command, args []string) {
		// https://iterm2.com/documentation-badges.html
		controlSequence := `"\e]1337;SetBadgeFormat=\a"`
		commandString := fmt.Sprintf("%s %s", "printf", controlSequence)

		cm := exec.Command(`bash`, `-c`, commandString)
		cm.Stdout = os.Stdout
		cm.Run()
	},
}

func init() {
	rootCmd.AddCommand(bclearCmd)
}

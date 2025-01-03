package commands

import (
	"fmt"
	"strings"
	"os/exec"

	"github.com/spf13/cobra"

	"shell-utils/ai"
)
  
func init() {
	rootCmd.AddCommand(saCmd)
}
  
var saCmd = &cobra.Command{
	Use:   "s",
	Short: "Shell assistant which generates shell command for your task. E.g. 'ai sa find all go files in current directory'",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		answer := ai.DoShellAssistant(query)
		Preview(answer.Explanation)
		fmt.Println(answer.Command)

		fields := strings.Fields(answer.Command)
		command := fields[0]
		cmdArgs := []string{}
		for _, field := range fields[1:] {
			// remove the single quotes around the field
			if strings.HasPrefix(field, "'") && strings.HasSuffix(field, "'") {
				cmdArgs = append(cmdArgs, field[1 : len(field)-1])
			} else {
				cmdArgs = append(cmdArgs, field)
			}
		}

		fmt.Println("Execute the command? (y/n)")
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "y" {
			return
		}

		osCmd := exec.Command(command, cmdArgs...)
		output, err := osCmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(output))
		}
	},
}
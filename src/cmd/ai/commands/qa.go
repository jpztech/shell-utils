package commands

import (
	"github.com/spf13/cobra"

	"shell-utils/ai"
)
  
func init() {
	rootCmd.AddCommand(qaCmd)
}
  
var qaCmd = &cobra.Command{
	Use:   "q [target]",
	Short: "Quick QA",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		answer := ai.QA(args)
		cmd.Println(answer)
	},
}
package commands

import (
	"strings"

	"github.com/spf13/cobra"

	"shell-utils/ai"
)
  
func init() {
	rootCmd.AddCommand(quickQaCmd)
}
  
var quickQaCmd = &cobra.Command{
	Use:   "q [your question]",
	Short: "Quick QA as a single-line command. E.g. ai q What is the capital of China?",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		answer := ai.QA(query)
		Preview(answer)
	},
}
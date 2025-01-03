package commands

import (
	"github.com/spf13/cobra"

	"shell-utils/ai"
	"shell-utils/utils"
)
  
func init() {
	rootCmd.AddCommand(qaCmd)
}
  
var qaCmd = &cobra.Command{
	Use:   "qa",
	Short: "General QA which supports multi-line input (type EOF to finish the input)",
	Args: cobra.MatchAll(cobra.MaximumNArgs(0), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		query := utils.GetMultiLineInput("What's your question?")
		answer := ai.DoQA(query)
		Preview(answer)
	},
}
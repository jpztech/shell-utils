package commands

import (
	"github.com/spf13/cobra"

	"shell-utils/ugit"
)
  
func init() {
	rootCmd.AddCommand(commitpushCmd)
}
  
var commitpushCmd = &cobra.Command{
	Use:   "commitpush <message>",
	Short: "Commit and push the changes to the current branch",
	Long: ugit.HELP_Commitpush,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	  ugit.Commitpush(args)
	},
}
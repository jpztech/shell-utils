package commands

import (
	"github.com/spf13/cobra"

	"shell-utils/ugit"
)
  
func init() {
	rootCmd.AddCommand(forcedefaultCmd)
}
  
var forcedefaultCmd = &cobra.Command{
	Use:   "fd",
	Short: "Force going back to the default branch (main/master) and delete the current branch",
	Args: cobra.MatchAll(cobra.ExactArgs(0), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	  ugit.Forcedefault()
	},
}
package commands

import (
	"github.com/spf13/cobra"

	"shell-utils/ugit"
)
  
func init() {
	rootCmd.AddCommand(branchesCmd)
}
  
var branchesCmd = &cobra.Command{
	Use:   "branches [target]",
	Short: "List the current branch for each repository",
	Long: ugit.HELP_Branches,
	Args: cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	  ugit.Branches(args)
	},
}
package commands

import (
	"github.com/spf13/cobra"

	"shell-utils/ugit"
)
  
func init() {
	rootCmd.AddCommand(cloneCmd)
}
  
var cloneCmd = &cobra.Command{
	Use:   "clone [target]",
	Short: "Clone a remote repository",
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	  ugit.Clone(args)
	},
}
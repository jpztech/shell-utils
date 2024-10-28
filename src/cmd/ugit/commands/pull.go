package commands

import (
	"github.com/spf13/cobra"

	"shell-utils/ugit"
)
  
func init() {
	rootCmd.AddCommand(pullCmd)
}
  
var pullCmd = &cobra.Command{
	Use:   "pull <target>",
	Short: "Pull the latest changes from the remote repository",
	Long: ugit.HELP_Pull,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
	  ugit.Pull(args)
	},
}
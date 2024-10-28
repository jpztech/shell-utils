package commands

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/spf13/cobra"

	"shell-utils/ugit"
)

var rootCmd = &cobra.Command{
	Use:   "ugit",
	Short: "ugit is a utility on top of git CLI to make git operations easier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Site: %s\n", Site)
	},
}

var Site string
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&Site, "site", "s", "GitHub", "GitHub site")
}

func initConfig() {
	binPath, err := os.Executable()
    if err != nil {
		fmt.Printf("Error getting executable folder: %v\n", err)
		os.Exit(1)
    }
	configFile := filepath.Join(filepath.Dir(binPath), "config.yaml")
	ugit.ReadConfigFromFile(configFile)
	err = ugit.ApplySite(Site)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
}
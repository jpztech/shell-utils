package commands

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/spf13/cobra"

	"shell-utils/ai"
)

var rootCmd = &cobra.Command{
	Use:   "ai",
	Short: "ai is a utility leverages GenAI for improving productivity",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	binPath, err := os.Executable()
    if err != nil {
		fmt.Printf("Error getting executable folder: %v\n", err)
		os.Exit(1)
    }
	configFile := filepath.Join(filepath.Dir(binPath), "ai.config.yaml")
	ai.ReadConfigFromFile(configFile)
	
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
}
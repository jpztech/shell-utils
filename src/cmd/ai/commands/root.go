package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"shell-utils/ai"
	"shell-utils/viewer/markdown"
)

var rootCmd = &cobra.Command{
	Use:   "ai",
	Short: "ai is a utility leverages GenAI for improving productivity",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// preview output as markdown
var OMarkdown bool
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&OMarkdown, "markdown", "m", true, "Preview output as markdown")
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

func Preview(input string) {
	if OMarkdown {
		output, err := markdown.Render(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println(input)
		} else {
			fmt.Println(output)
		}
	} else {
		fmt.Println(input)
	}
}
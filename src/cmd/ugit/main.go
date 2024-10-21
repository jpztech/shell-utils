package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"shell-utils/ugit"
)

func main() {
	binPath, err1 := os.Executable()
    if err1 != nil {
		fmt.Printf("Error getting executable folder: %v\n", err1)
		os.Exit(1)
    }
	configFile := filepath.Join(filepath.Dir(binPath), "config.yaml")
	site := flag.String("site", "GitHub", "GitHub site")
	flag.Parse()
	args := flag.Args()

	config := ugit.FromFile(configFile)
	_, err := config.ApplySite(*site)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(args) == 0 {
		fmt.Println("Please specify a command")
		os.Exit(1)
	}
	ugit.Run(config, args)
}
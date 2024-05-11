package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// Executes a command and returns the output
func ExecCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing command %s: %v\n", command, err)
		os.Exit(1)
	}
	fmt.Println(string(output))
	return string(output), nil
}
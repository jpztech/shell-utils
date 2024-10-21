package utils

import (
	"os"
	"os/exec"

	"github.com/fatih/color"	
)

func ExecCommand(command string, args ...string) (string, error) {
	return execCommand(true, nil, command, args...)
}

func ExecCommandIn(path *string, command string, args ...string) (string, error) {
	return execCommand(false, path, command, args...)
}

func ExecCommandSilentIn(path *string, command string, args ...string) (string, error) {
	return execCommand(true, path, command, args...)
}

// Executes a command and returns the output
func execCommand(silent bool, path *string, command string, args ...string) (string, error) {
	if !silent {
		color.Green("Executing command: %s %v\n\n", command, args)
	}
	cmd := exec.Command(command, args...)
	if !silent {
		cmd.Stderr = os.Stderr
	}
	if path != nil {
		cmd.Dir = *path
	}
	output, err := cmd.Output()
	if !silent && err != nil {
		color.Red("Error executing command %s: %v\n", command, err)
		os.Exit(1)
	}
	return string(output), err
}
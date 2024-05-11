package command

import (
	"fmt"
	"os"

	"shell-utils/git"
)

type Command struct {
	Name 		string
	ArgsLength	int
	Func 		func(args []string)
	Prechecks 	[]func(args []string) bool
}

type Commands struct {
	Commands []Command
}

// Check whether args length is matching the ArgsLength of the command and then call Func
func (c *Command) Call(args []string) {
	if len(args) != c.ArgsLength {
		fmt.Printf("Unexpected arguments for %s, require %d args\n", c.Name, c.ArgsLength)
		os.Exit(1)
	}
	c.Func(args)
}

// Dispatch to the appropriate Command according to the first argument as the command name
func (c *Commands) Dispatch(args []string) {
	for _, command := range c.Commands {
		if args[0] == command.Name {
			command.Call(args[1:])
			return
		}
	}
	fmt.Printf("Unknown command: %s\n", args[0])
	os.Exit(1)
}

// Check whether the current directory is a git repository
func IsGitRepo(args []string) bool {
	if git.IsGitRepo(".") {
		return true
	}
	fmt.Println("Current directory is not a git repository")
	return false
}
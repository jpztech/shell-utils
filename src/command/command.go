package command

import (
	"fmt"
	"os"
	"strconv"

	"shell-utils/git"
)

type Command struct {
	Name 		string
	ArgsLen		string
	Func 		func(args []string)
	Prechecks 	[]func(args []string) bool
}

type Commands struct {
	Commands []Command
}

// Check whether args length is matching the ArgsLength of the command and then call Func
func (c *Command) Call(args []string) {
	mode := "eq"
	argsLenStr := c.ArgsLen
	// if ArgsLen ends with +, it means the args length should be greater than or equal to ArgsLen
	if c.ArgsLen[len(c.ArgsLen)-1] == '+' {
		mode = "ge"
		argsLenStr = c.ArgsLen[:len(c.ArgsLen)-1]
	}
	if argsLen, err := strconv.Atoi(argsLenStr); err != nil {
		ok := len(args) == argsLen
		if mode == "ge" {
			ok = len(args) >= argsLen
		}
		if !ok {
			fmt.Printf("Unexpected arguments for %s, require %s args\n", c.Name, c.ArgsLen)
			os.Exit(1)
		}
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

// A command to check whether the current directory is a git repository
func IsGitRepo(args []string) bool {
	if git.IsGitRepo(".") {
		return true
	}
	fmt.Println("Current directory is not a git repository")
	return false
}
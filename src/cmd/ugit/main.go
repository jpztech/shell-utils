package main

import (
	"fmt"
	"os"

	"shell-utils/command"
	"shell-utils/git"
	"shell-utils/utils"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please specify a command")
		os.Exit(1)
	}
	requireGitRepo := []func(args []string) bool{command.IsGitRepo}
	commands := command.Commands{
		Commands: []command.Command{
			{Name: "forcemain", ArgsLength: 0, Func: forcemain, Prechecks: requireGitRepo},
			{Name: "commitpush", ArgsLength: 1, Func: commitpush, Prechecks: requireGitRepo},
		},
	}
	commands.Dispatch(args)
}

// Forcing going back to main branch and pulling and delete the branch going from
func forcemain(args []string) {
	repo := "."
	
	curBranch := git.CurrentBranch(repo)
	if git.CurrentBranch(repo) == "main" {
		fmt.Println("Already on main branch")
		os.Exit(0)
	}
	utils.ExecCommand("git", "checkout", "main")
	utils.ExecCommand("git", "pull", "origin", "main")
	utils.ExecCommand("git", "branch", "-D", curBranch)
}

// Commit and push the changes to the current branch
func commitpush(args []string) {
	repo := "."
	commitMsg := args[0]
	
	if git.CurrentBranch(repo) == "main" {
		fmt.Println("Cannot commit to main branch")
		os.Exit(1)
	}
	utils.ExecCommand("git", "add", ".")
	utils.ExecCommand("git", "commit", "-m", commitMsg)
	utils.ExecCommand("git", "push", "origin", git.CurrentBranch(repo))
}
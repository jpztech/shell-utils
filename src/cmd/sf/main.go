package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"shell-utils/git"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Please specify a command")
		os.Exit(1)
	}
	if args[0] == "clone" {
		if len(args) != 2 {
			fmt.Println("Please specify a repository")
			os.Exit(1)
		}
		clone(args[1])
	} else if args[0] == "pull" {
		pull(args[1:])
	} else if args[0] == "branch" {
		branches(args[1:])
	} else {
		fmt.Printf("Unknown command: %s\n", args[0])
		os.Exit(1)
	}
}

func clone(repo string) {
	cmd := exec.Command("git", "clone", repo)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error cloning repository:", err)
		os.Exit(1)
	}
}

func pull(args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify a repository")
		return
	}
	repos := args
	if len(args) == 1 {
		if args[0] == "all" {
			repos = getAllRepos(nil)
		} else if !git.IsGitRepo(args[0]) {
			repos = getAllRepos(&args[0])
		}
	}
	fmt.Printf("Repos found: %v\n", repos)
	for _, repo := range repos {
		fmt.Printf("Pulling %s...\n", repo)
		git.PullBranch(repo, "master")
	}
}

// branches lists current branch for all repositories
func branches(args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify a repository")
		return
	}
	repos := args
	if len(args) == 1 {
		if args[0] == "all" {
			repos = getAllRepos(nil)
		} else if !git.IsGitRepo(args[0]) {
			repos = getAllRepos(&args[0])
		}
	}
	fmt.Printf("Repos found: %v\n", repos)
	for _, repo := range repos {
		branch := git.CurrentBranch(repo)
		if branch != "" {
			fmt.Printf("%s: %s\n", repo, branch)
		}
	}
}

// getAllRepos returns a list of all repositories in the current directory
func getAllRepos(pattern *string) []string {
	entries, err := os.ReadDir("./")
	if err != nil {
		fmt.Println("Error listing repositories:", err)
		os.Exit(1)
	}
	var repos []string
	for _, repo := range entries {
		if repo.IsDir() && git.IsGitRepo(repo.Name()) {
			if pattern != nil && !strings.Contains(repo.Name(), *pattern) {
				continue
			}
			repos = append(repos, repo.Name())
		}
	}
	return repos
}
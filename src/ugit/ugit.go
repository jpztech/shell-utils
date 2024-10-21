package ugit

import (
	"fmt"
	"os"
	"shell-utils/command"
	"shell-utils/git"
	"shell-utils/utils"
	"strings"

	"github.com/fatih/color"
)

var config *Config

func Run(_config *Config, args []string) {
	config = _config
	requireGitRepo := []func(args []string) bool{command.IsGitRepo}
	noCheck := []func(args []string) bool{}
	commands := command.Commands{
		Commands: []command.Command{
			{Name: "clone", ArgsLen: "1", Func: clone, Prechecks: noCheck},
			{Name: "pull", ArgsLen: "1+", Func: pull, Prechecks: noCheck},
			{Name: "branches", ArgsLen: "1+", Func: branches, Prechecks: noCheck},
			{Name: "forcemain", ArgsLen: "0", Func: forcemain, Prechecks: requireGitRepo},
			{Name: "forcemaster", ArgsLen: "0", Func: forcemaster, Prechecks: requireGitRepo},
			{Name: "commitpush", ArgsLen: "1", Func: commitpush, Prechecks: requireGitRepo},
		},
	}
	commands.Dispatch(args);
}

func forcemain(args []string) {
	forcedefault("main")
}
func forcemaster(args []string) {
	forcedefault("master")
}

// Forcing going back to default branch and pulling and delete the branch going from
func forcedefault(bname string) {
	repo := "."
	
	curBranch := git.CurrentBranch(repo)
	if git.CurrentBranch(repo) == bname {
		fmt.Printf("Already on %s branch\n", bname)
		os.Exit(0)
	}
	utils.ExecCommand("git", "checkout", bname)
	utils.ExecCommand("git", "pull", "origin", bname)
	utils.ExecCommand("git", "branch", "-D", curBranch)
}

// Commit and push the changes to the current branch
func commitpush(args []string) {
	repo := "."
	commitMsg := args[0]
	
	if git.CurrentBranch(repo) == config.CurrentSite.DefaultBranch {
		color.Red("Cannot commit to default branch %s\n", config.CurrentSite.DefaultBranch)
		os.Exit(1)
	}
	utils.ExecCommand("git", "add", ".")
	utils.ExecCommand("git", "commit", "-m", commitMsg)
	utils.ExecCommand("git", "push", "origin", git.CurrentBranch(repo))
}

func clone(args []string) {
	utils.ExecCommand("git", "clone", config.GetURL(args[0]))
}

// pull all: pulls all repositories in the current directory
// pull <repo>: pulls the specified repository
// pull <pattern>: pulls all repositories matching the pattern
func pull(args []string) {
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
		git.PullBranch(repo, config.CurrentSite.DefaultBranch)
	}
}

// branches: lists current branch for all repositories
// branches all: lists current branch for all repositories
// branches <repo>: lists current branch for the specified repository
// branches <pattern>: lists current branch for all repositories matching the pattern
func branches(args []string) {
	repos := args
	if len(args) == 0 {
		repos = getAllRepos(nil)
	}
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
			fmt.Printf("%s: ", repo)
			color.Green("%s\n", branch)
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
		if !strings.HasPrefix(repo.Name(), ".") && repo.IsDir() && git.IsGitRepo(repo.Name()) {
			if pattern != nil && !strings.Contains(repo.Name(), *pattern) {
				continue
			}
			repos = append(repos, repo.Name())
		}
	}
	return repos
}
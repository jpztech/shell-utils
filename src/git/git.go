package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Returns true if the given path is a git repository
func IsGitRepo(path string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = path
	err := cmd.Run()
	return err == nil
}

// Returns the current branch for the given repository
func CurrentBranch(repo string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = repo
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting current branch for %s: %v\n", repo, err)
		return ""
	}
	return strings.TrimSpace(string(output))
}


// Pulls the specified branch from the given repository
func PullBranch(repo string, branch string) {
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Dir = repo
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error pulling branch %s from %s: %v\n", branch, repo, err)
	}
}
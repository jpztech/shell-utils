package git

import (
	"strings"

	"github.com/fatih/color"	

	"shell-utils/utils"
)

// Returns true if the given path is a git repository
func IsGitRepo(path string) bool {
	_, err := utils.ExecCommandSilentIn(&path, "git", "rev-parse", "--is-inside-work-tree")
	return err == nil
}

// Returns the current branch for the given repository
func CurrentBranch(repo string) string {
	output, _ := utils.ExecCommandSilentIn(&repo, "git", "rev-parse", "--abbrev-ref", "HEAD")
	return strings.TrimSpace(string(output))
}


// Pulls the specified branch from the given repository
func PullBranch(repo string, branch string) {
	_, err := utils.ExecCommandIn(&repo, "git", "pull", "origin", branch)
	if err != nil {
		color.Red("Error pulling branch %s from %s: %v\n", branch, repo, err)
	}
}
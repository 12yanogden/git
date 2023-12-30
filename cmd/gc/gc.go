package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/12yanogden/cli"
	"github.com/12yanogden/git/internal/git"
	"github.com/12yanogden/git/internal/ticket"
	"github.com/12yanogden/shell"
)

func main() {
	args := os.Args[1:]
	branch := git.CurrentBranch()
	ticket := ticket.BranchToTicket(branch)

	if len(args) == 0 {
		fmt.Println("gc: commit message required")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Printf("gc: expected 1 argument. %d arguments given", len(args))
	}

	if strings.Contains(branch, "develop") ||
		strings.Contains(branch, "master") ||
		strings.Contains(branch, "epic/") {
		msg := "WARNING:\nBranch " + branch + " might be sensitive. Are you sure you want to commit? (y/n)"

		if !cli.Confirm(msg) {
			fmt.Println("No changes were added, committed, or pushed")
			os.Exit(1)
		}
	}

	shell.Run("git", []string{"add", "."})
	shell.Run("git", []string{"commit", "-m", ticket + args[0]})
	shell.Run("git", []string{"push"})
}

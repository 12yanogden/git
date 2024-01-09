package main

import (
	"os"
	"strings"

	"github.com/12yanogden/git/internal/git"
	"github.com/12yanogden/git/internal/ticket"
	"github.com/12yanogden/shell"
	"github.com/pterm/pterm"
)

func main() {
	args := os.Args[1:]
	branch := git.CurrentBranch()
	ticket := ticket.BranchToTicket(branch)

	if len(args) == 0 {
		pterm.Error.Println("gc: commit message required")
		os.Exit(1)
	} else if len(args) > 1 {
		pterm.Error.Printf("gc: expected 1 argument. %d arguments given", len(args))
		os.Exit(1)
	}

	if strings.Contains(branch, "master") ||
		strings.Contains(branch, "main") ||
		strings.Contains(branch, "dev") ||
		strings.Contains(branch, "epic/") {

		pterm.Warning.Printf("\nYou are committing directly to %s ", branch)

		pterm.Println()
		pterm.Println()

		isConfirmed, _ := pterm.DefaultInteractiveConfirm.Show()

		pterm.Println()

		if !isConfirmed {
			pterm.Info.Println("\nNo changes were added, committed, or pushed\n ")
			pterm.Println()
			os.Exit(1)
		}
	}

	multi := pterm.DefaultMultiPrinter

	addSpinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Add changes")
	commitSpinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Add changes")
	pushSpinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start("Add changes")

	multi.Start()

	shell.Run("git", []string{"add", "."})

	addSpinner.Success()

	shell.Run("git", []string{"commit", "-m", ticket + args[0]})

	commitSpinner.Success()

	shell.Run("git", []string{"push"})

	pushSpinner.Success()
}

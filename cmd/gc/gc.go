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
	ticket := ticket.BranchToTicket(&branch)

	if len(args) == 0 {
		pterm.Error.Println("gc: commit message required")
		os.Exit(1)
	} else if len(args) > 1 {
		pterm.Error.Printf("gc: expected 1 argument. %d arguments given", len(args))
		os.Exit(1)
	}

	if len(ticket) > 0 {
		ticket += ": "
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

	addMsg := "Add changes"
	commitMsg := "Commit changes"
	pushMsg := "Push changes"

	multi := pterm.DefaultMultiPrinter

	addSpinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start(addMsg)
	commitSpinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start(commitMsg)
	pushSpinner, _ := pterm.DefaultSpinner.WithWriter(multi.NewWriter()).Start(pushMsg)

	multi.Start()

	shell.Run("git", []string{"add", "."})

	addSpinner.Stop()
	pterm.Println("[" + pterm.Green("✓") + "] " + addMsg)

	shell.Run("git", []string{"commit", "-m", ticket + args[0]})

	commitSpinner.Stop()
	pterm.Println("[" + pterm.Green("✓") + "] " + commitMsg)

	shell.Run("git", []string{"push"})
	pterm.Println("[" + pterm.Green("✓") + "] " + pushMsg)

	pushSpinner.Stop()
}

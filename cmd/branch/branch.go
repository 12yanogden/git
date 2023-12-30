package branch

import (
	"fmt"
	"os"

	"github.com/12yanogden/git/internal/git"
	"github.com/12yanogden/shell"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printCurrentBranch()
	} else {
		printMatchingBranches(args)
	}

	os.Exit(0)
}

func printCurrentBranch() {
	out := git.CurrentBranch()

	if shell.IsTerminal() {
		out += "\n"
	}

	fmt.Printf("%s", out)
}

func printMatchingBranches(patterns []string) {
	branches := git.MatchingBranches(patterns)

	for _, branch := range branches {
		fmt.Println(branch)
	}
}

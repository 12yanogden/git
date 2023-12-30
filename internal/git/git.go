package git

import (
	"strings"

	"github.com/12yanogden/shell"
	islices "github.com/12yanogden/slices"
	"github.com/12yanogden/str"
	"github.com/12yanogden/tables"
)

type Branch struct {
	Name   string
	Head   string
	Remote string
}

func CurrentBranch() string {
	return shell.Run("git", []string{"rev-parse", "--abbrev-ref", "HEAD"})
}

func AllBranches() []Branch {
	branchData := getBranchData()
	branches := []Branch{}

	for _, branchDatum := range branchData {
		if isRemote(branchDatum[0]) {
			branches = append(branches,
				Branch{
					Name:   strings.TrimPrefix(branchDatum[0], "remotes/"),
					Head:   branchDatum[1],
					Remote: "",
				},
			)
		} else {
			branches = append(branches,
				Branch{
					Name:   branchDatum[0],
					Head:   branchDatum[1],
					Remote: str.TrimSides(branchDatum[2], 1, 1),
				},
			)
		}
	}

	return branches
}

func isRemote(branch string) bool {
	return strings.Contains(branch, "remotes/")
}

func getBranchData() [][]string {
	slices := tables.Slices(shell.Run("git", []string{"branch", "-vva"}))

	slices = islices.RemoveEmpty(slices)

	slices = rejectCommits(slices)

	for i, slice := range slices {
		if slice[0] == "*" {
			slices[i] = slice[1:]
		}
	}

	return slices
}

func rejectCommits(slices [][]string) [][]string {
	commitIndexes := []int{}

	// Collect rows with commits
	for i, slice := range slices {
		if slice[1] == "->" {
			commitIndexes = islices.Prepend(commitIndexes, i)
		}
	}

	// Remove rows with commits
	for _, index := range commitIndexes {
		slices = islices.Remove[[]string](slices, index)
	}

	return slices
}

func MatchingBranches(patterns []string) []string {
	branches := []string{}

	for _, pattern := range patterns {
		allBranches := AllBranches()

		for _, branch := range allBranches {
			if strings.Contains(branch.Name, pattern) {
				branches = append(branches, branch.Name)
			}
		}
	}

	return branches
}

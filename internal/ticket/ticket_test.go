package ticket

import (
	"testing"
)

func TestBranchToTicket(t *testing.T) {
	branch := "BPM-2893-this-that"
	expected := "BPM-2893"
	actual := BranchToTicket(branch)

	if expected != actual {
		t.Fatalf("\nExpected:\t'" + expected + "'\nActual:\t\t'" + actual + "'")
	}
}

func TestBranchWithoutTicketToBlank(t *testing.T) {
	branch := "-this-that"
	expected := ""
	actual := BranchToTicket(branch)

	if expected != actual {
		t.Fatalf("\nExpected:\t'" + expected + "'\nActual:\t\t'" + actual + "'")
	}
}

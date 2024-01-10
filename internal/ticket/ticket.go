package ticket

import "regexp"

func BranchToTicket(branch *string) string {
	return string(regexp.MustCompile(`^[A-Z]+\-[0-9]+`).Find([]byte(*branch)))
}

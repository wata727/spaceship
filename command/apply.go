package command

import (
	"strings"
)

type ApplyCommand struct {
	Meta
}

func (c *ApplyCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *ApplyCommand) Synopsis() string {
	return ""
}

func (c *ApplyCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

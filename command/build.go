package command

import (
	"strings"
)

type BuildCommand struct {
	Meta
}

func (c *BuildCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *BuildCommand) Synopsis() string {
	return ""
}

func (c *BuildCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

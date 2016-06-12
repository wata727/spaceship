package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type InitCommand struct {
	Meta
}

func (c *InitCommand) Run(args []string) int {
	// Write your code here

	if err := os.Mkdir(".spaceship", 0755); err != nil {
		c.Ui.Error("Failed initialization: " + err.Error())
		return 1
	}
	if err := os.Mkdir("artifact", 0755); err != nil {
		c.Ui.Error("Failed initialization: " + err.Error())
		return 1
	}
	if err := os.Mkdir("template", 0755); err != nil {
		c.Ui.Error("Failed initialization: " + err.Error())
		return 1
	}
	if err := os.Mkdir("infra", 0755); err != nil {
		c.Ui.Error("Failed initialization: " + err.Error())
		return 1
	}
	if err := os.Mkdir("strategy", 0755); err != nil {
		c.Ui.Error("Failed initialization: " + err.Error())
		return 1
	}
	var git_ignore = fmt.Sprintf(`
/infra/spaceship-variables.tf
`)
	ioutil.WriteFile(".gitignore", []byte(git_ignore), os.ModePerm)

	c.Ui.Output("Initialized empty spaceship project!")

	return 0
}

func (c *InitCommand) Synopsis() string {
	return "Initialize spaceship project"
}

func (c *InitCommand) Help() string {
	helpText := `
Initialize spaceship project
`
	return strings.TrimSpace(helpText)
}

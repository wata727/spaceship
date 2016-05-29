package command

import (
	"os"
	"strings"

	client "github.com/wata727/spaceship/packer"
	"github.com/wata727/spaceship/vui"
)

type BuildCommand struct {
	Meta
}

func (c *BuildCommand) Run(args []string) int {
	// Write your code here

	pwd, _ := os.Getwd()
	packer := &client.Packer{
		Dir:      pwd + "/template",
		Varfile:  "variables.json",
		Template: "template.json",
		Ui:       vui.Ui{},
	}

	c.Ui.Output("Building Artifact...")
	if err := packer.Build(); err != nil {
		c.Ui.Error("Failed packer build: " + err.Error())
	}
	c.Ui.Output(packer.Artifact.Output())
	packer.Ui.Flush()

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

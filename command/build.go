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
		Artifact: client.Artifact{
			Dir: pwd + "/.spaceship",
		},
		Ui: vui.Ui{},
	}
	packer.Artifact.Init()

	c.Ui.Output("Building Artifact...")
	if err := packer.Build(); err != nil {
		c.Ui.Error("Failed packer build: " + err.Error())
	}
	if err := packer.Artifact.Save(); err != nil {
		c.Ui.Error("Failed save artifact: " + err.Error())
	}
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

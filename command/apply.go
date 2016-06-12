package command

import (
	"os"
	"strings"

	"github.com/wata727/spaceship/packer"
	client "github.com/wata727/spaceship/terraform"
	"github.com/wata727/spaceship/vui"
)

type ApplyCommand struct {
	Meta
}

func (c *ApplyCommand) Run(args []string) int {
	// Write your code here

	pwd, _ := os.Getwd()
	terraform := &client.Terraform{
		Dir: pwd + "/infra",
		Artifact: packer.Artifact{
			Dir: pwd + "/artifact",
		},
		Ui: vui.Ui{},
	}

	c.Ui.Output("Applying Infrastructure...")
	if err := terraform.Apply(); err != nil {
		c.Ui.Error("Failed terraform apply: " + err.Error())
	}
	terraform.Ui.Flush()

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

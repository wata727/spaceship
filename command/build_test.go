package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestBuildCommand_implement(t *testing.T) {
	var _ cli.Command = &BuildCommand{}
}

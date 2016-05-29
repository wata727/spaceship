package packer

import "strings"

type Artifact struct {
	artifact string
}

func (a *Artifact) Detect(output string) error {
	parts := strings.Split(output, ",")
	if len(parts) < 5 || parts[4] != "id" {
		return nil
	}
	artifact := strings.Split(parts[5], ":")
	a.artifact = artifact[1]
	return nil
}

func (a *Artifact) Output() string {
	return a.artifact
}

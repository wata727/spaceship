package packer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Artifact struct {
	Name string `json:"name"`
	Dir  string `json:"-"`
}

func (a *Artifact) Detect(output string) error {
	parts := strings.Split(output, ",")
	if len(parts) < 5 || parts[4] != "id" {
		return nil
	}
	artifact := strings.Split(parts[5], ":")
	name := strings.Replace(artifact[1], "\n", "", -1)
	name = strings.Replace(name, "\r", "", -1)
	a.Name = name
	return nil
}

func (a *Artifact) Save() error {
	artifact_file := a.Dir + "/artifact.json"
	artifact_body, err := json.Marshal(a)
	if err != nil {
		return err
	}
	ioutil.WriteFile(artifact_file, artifact_body, os.ModePerm)
	return nil
}

func (a *Artifact) Parse(workdir string) error {
	artifact_file := a.Dir + "/artifact.json"
	variables_file := workdir + "/spaceship_variables.tf"
	artifact_body, err := ioutil.ReadFile(artifact_file)
	if err != nil {
		return err
	}

	json.Unmarshal(artifact_body, &a)

	var tfvars_body = fmt.Sprintf(`
variable "%s" {
    default = "%s"
}
`, "spaceship_artifact", a.Name)
	ioutil.WriteFile(variables_file, []byte(tfvars_body), os.ModePerm)

	return nil
}

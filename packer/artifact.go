package packer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Artifact struct {
	Name []string `json:"name"`
	Dir  string   `json:"-"`
}

func (a *Artifact) Init() error {
	artifact_file := a.Dir + "/artifact.json"
	artifact_body, err := ioutil.ReadFile(artifact_file)
	if err != nil {
		return err
	}

	json.Unmarshal(artifact_body, &a)

	return nil
}

func (a *Artifact) Detect(output string) error {
	parts := strings.Split(output, ",")
	if len(parts) < 5 || parts[4] != "id" {
		return nil
	}
	artifact := strings.Split(parts[5], ":")
	name := strings.Replace(artifact[1], "\n", "", -1)
	name = strings.Replace(name, "\r", "", -1)
	if len(a.Name) > 4 {
		old_image := a.Name[0]
		a.Name = a.Name[1:5]
		if err := a.delete(old_image); err != nil {
			return err
		}
	}
	a.Name = append(a.Name, name)

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
`, "spaceship_artifact", a.Name[len(a.Name)-1])
	ioutil.WriteFile(variables_file, []byte(tfvars_body), os.ModePerm)

	return nil
}

func (a *Artifact) delete(name string) error {
	aws_client := ec2.New(session.New())

	output, err := aws_client.DescribeImages(&ec2.DescribeImagesInput{
		ImageIds: aws.StringSlice([]string{name}),
	})
	if err != nil {
		return err
	}

	if _, err := aws_client.DeregisterImage(&ec2.DeregisterImageInput{
		ImageId: output.Images[0].ImageId,
	}); err != nil {
		return err
	}

	for _, device := range output.Images[0].BlockDeviceMappings {
		if _, err := aws_client.DeleteSnapshot(&ec2.DeleteSnapshotInput{
			SnapshotId: device.Ebs.SnapshotId,
		}); err != nil {
			return err
		}
	}

	return nil
}

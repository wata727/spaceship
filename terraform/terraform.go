package terraform

import (
	"io"
	"os"
	"os/exec"

	"github.com/wata727/spaceship/packer"
	"github.com/wata727/spaceship/vui"
)

type Terraform struct {
	Dir      string
	Artifact packer.Artifact
	Ui       vui.Ui
}

func (t *Terraform) Apply() error {
	terraform_r, terraform_w := io.Pipe()
	if err := t.Artifact.Parse(t.Dir); err != nil {
		return err
	}

	reader := make(chan bool)
	go func() {
		defer close(reader)
		var buf [1024]byte
		for {
			n, err := terraform_r.Read(buf[:])
			if err != nil {
				break
			}
			if n > 0 {
				t.Ui.Output(string(buf[:n]))
			}
		}
	}()

	terraform_arg := []string{"apply"}

	cmd := exec.Command("terraform", terraform_arg...)
	cmd.Stdout = io.MultiWriter(os.Stdout, terraform_w)
	cmd.Stderr = io.MultiWriter(os.Stderr, terraform_w)
	cmd.Dir = t.Dir

	cmderr := cmd.Run()
	terraform_w.Close()
	<-reader
	if cmderr != nil {
		return cmderr
	}

	return nil
}

package packer

import (
	"io"
	"os"
	"os/exec"

	"github.com/wata727/spaceship/vui"
)

type Packer struct {
	Dir      string
	Varfile  string
	Template string
	Ui       vui.Ui
	Artifact Artifact
}

func (p *Packer) Build() error {
	packer_r, packer_w := io.Pipe()

	reader := make(chan error)
	go func() {
		defer close(reader)
		var buf [1024]byte
		for {
			n, err := packer_r.Read(buf[:])
			if err != nil {
				break
			}
			if n > 0 {
				p.Ui.Output(string(buf[:n]))
				if err := p.Artifact.Detect(string(buf[:n])); err != nil {
					reader <- err
				}
			}
		}
		reader <- nil
	}()

	packer_arg := []string{"build"}
	if p.Varfile != "" {
		packer_arg = append(packer_arg, "-var-file=")
		packer_arg = append(packer_arg, p.Varfile)
	}
	packer_arg = append(packer_arg, "-machine-readable")
	packer_arg = append(packer_arg, p.Template)

	cmd := exec.Command("packer", packer_arg...)
	cmd.Stdout = io.MultiWriter(os.Stdout, packer_w)
	cmd.Stderr = io.MultiWriter(os.Stderr, packer_w)
	cmd.Dir = p.Dir

	err := cmd.Run()
	packer_w.Close()
	err_chan := <-reader
	if err != nil {
		return err
	}
	if err_chan != nil {
		return err_chan
	}

	return nil
}

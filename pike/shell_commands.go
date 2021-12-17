package pike

import (
	"log"
	"os/exec"
	"path/filepath"
)

func (p Project) compileProto() (err error) {
	_, err = exec.LookPath("protoc")
	if err != nil {
		return
	}

	cmdPath := filepath.Join(p.AbsolutePath(), "bin/compile_proto.sh")
	cmd := exec.Command(cmdPath)
	err = cmd.Run()
	return
}

func (p Project) RunGoModInit() {
	cmd := exec.Command("go", "mod", "init")
	cmd.Dir = p.AbsolutePath()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing go mod init: %v", err)
	}
}

func (p Project) RunGoModTidy() {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = p.AbsolutePath()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing go mod init: %v", err)
	}
}

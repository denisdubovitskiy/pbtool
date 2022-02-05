package protoc

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Plugin struct {
	Name    string
	Path    string
	Options []string
}

func Compile(filePath string, vendorDir string, plugins []Plugin) error {
	cmd := []string{
		"--proto_path", filepath.Dir(filePath),
		"--proto_path", vendorDir,
	}

	for _, plug := range plugins {
		cmd = append(cmd, "--"+plug.Name+"_out", filepath.Dir(filePath))

		for _, opt := range plug.Options {
			cmd = append(cmd, "--"+plug.Name+"_opt", opt)
		}
	}

	cmd = append(cmd, filePath)

	c := exec.Command("protoc", cmd...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	log.Printf("builder: started building %s\n", filePath)
	command := strings.Join(append([]string{"protoc"}, cmd...), " ")
	command = strings.ReplaceAll(command, " --", " \\\n --")
	log.Println(command)

	if err := c.Run(); err != nil {
		log.Fatalf("unable to compile: %v", err)
	}

	return nil
}

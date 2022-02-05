package gitrepo

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Clone(repo, dst string) error {
	stat, err := os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("gitrepo: unable to stat a dir %s:%w", dst, err)
	}

	if stat != nil && stat.IsDir() {
		return nil
	}

	if err := os.MkdirAll(dst, 0755); err != nil {
		return fmt.Errorf("gitrepo: unable to create a dir %s:%w", dst, err)
	}

	log.Printf("cloning %s", repo)

	cmd := exec.Command("git", "clone", repo, dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("gitrepo: unable to clone a repo %s:%w", repo, err)
	}

	return nil
}

package filesystem

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

type FS struct {
	fs afero.Fs
}

func NewOsFs() *FS {
	return &FS{fs: afero.NewOsFs()}
}

func New(fs afero.Fs) *FS {
	return &FS{fs: fs}
}

func (f *FS) CreateDirectory(dir string) error {
	if err := f.fs.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("filesystem: unable to create a directory %s:%v", dir, err)
	}

	return nil
}

func (f *FS) CopyFile(src, dst string) error {
	ff, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("filesystem: unable to read a file %s: %v", src, err)
	}

	if err := os.WriteFile(dst, ff, 0644); err != nil {
		return fmt.Errorf("filesystem: unable to write a file %s: %v", dst, err)
	}

	return nil
}

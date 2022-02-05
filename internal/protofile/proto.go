package protofile

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/emicklei/proto"
)

type Finder struct{}

func NewFinder() *Finder {
	return &Finder{}
}

func (f *Finder) FindImports(reader io.Reader) ([]string, error) {
	parser := proto.NewParser(reader)
	ast, err := parser.Parse()
	if err != nil {
		return nil, fmt.Errorf("protofile: unable to parse %v", err)
	}

	var imports []string

	proto.Walk(ast,
		proto.WithImport(func(i *proto.Import) {
			imports = append(imports, i.Filename)
		}),
	)

	return imports, nil
}

func (f *Finder) FindImportedProtofilesInFile(path string) ([]string, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := fh.Close(); err != nil {
			log.Printf("protofile: unable to close a file %s: %v", path, err)
		}
	}()

	imports, err := f.FindImports(fh)
	if err != nil {
		return nil, fmt.Errorf("protofile: unable to find imports %s: %v", path, err)
	}

	return imports, nil
}

func (f *Finder) FindProtoFiles(dirs []string) []string {
	var protos []string

	for _, dir := range dirs {
		filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			if strings.HasSuffix(path, ".proto") {
				protos = append(protos, path)
			}
			return nil
		})
	}

	return protos
}

package vendorer

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/denisdubovitskiy/pbtool/internal/config"
	"github.com/denisdubovitskiy/pbtool/internal/depmanager"
	"github.com/denisdubovitskiy/pbtool/internal/filesystem"
	"github.com/denisdubovitskiy/pbtool/internal/gitrepo"
	"github.com/denisdubovitskiy/pbtool/internal/pathmanager"
	"github.com/denisdubovitskiy/pbtool/internal/protofile"
)

func OptsFromConfig(projConf *config.Project, repoCacheDir string) VedorOpts {
	return VedorOpts{
		Rules:           projConf.ReplaceRules,
		ExternalDeps:    projConf.ExternalDeps,
		ProtoDirs:       projConf.ProtoDirs,
		RepoCacheDir:    repoCacheDir,
		ProtoVendorDir:  projConf.ProtoVendorDir,
		ExternalDepsDir: projConf.ExternalDepsDir,
	}
}

type VedorOpts struct {
	Rules           []config.Rule
	ExternalDeps    []config.ExternalDependency
	ProtoDirs       []string
	ExternalDepsDir string
	RepoCacheDir    string
	ProtoVendorDir  string
}

func New(opts VedorOpts) *Vendorer {
	return &Vendorer{
		Rules:         opts.Rules,
		External:      opts.ExternalDeps,
		ProtoDirs:     opts.ProtoDirs,
		fs:            filesystem.NewOsFs(),
		pathFormatter: pathmanager.New(opts.RepoCacheDir, opts.ProtoVendorDir, opts.ExternalDepsDir),
		dependencies:  depmanager.NewDeduplicateDependencyManager(),
		gitrepo:       gitrepo.New(),
		finder:        protofile.NewFinder(),
	}
}

type Vendorer struct {
	ProtoDirs     []string
	Rules         []config.Rule
	External      []config.ExternalDependency
	fs            FileSystem
	dependencies  DependencyManager
	pathFormatter PathFormatter
	gitrepo       Cloner
	finder        Finder
}

func (v *Vendorer) VendorDependencies() error {
	for _, e := range v.External {
		if err := v.gitrepo.Clone(e.RepoSSHCloneURL(), v.pathFormatter.FormatCachePath(e.RepoDir())); err != nil {
			return fmt.Errorf("unable to clone repo: %v", err)
		}

		dstProtoFile := v.pathFormatter.FormatExternalPath(e.Filename())
		dstProtoFileDir := filepath.Dir(dstProtoFile)

		if err := v.fs.CreateDirectory(dstProtoFileDir); err != nil {
			return fmt.Errorf("unable to create directory: %s", err)
		}

		externalDepFilename := v.pathFormatter.FormatCachePath(e.Filename())

		if err := v.fs.CopyFile(externalDepFilename, dstProtoFile); err != nil {
			return fmt.Errorf("unable to copy an external dependency")
		}

		importedFiles, err := v.finder.FindImportedProtofilesInFile(dstProtoFile)
		if err != nil {
			return fmt.Errorf("unable to find dependencies in file %s:%v", dstProtoFile, err)
		}

		v.dependencies.EnqueueDependencies(importedFiles...)
	}

	for _, p := range v.finder.FindProtoFiles(v.ProtoDirs) {
		importedFiles, err := v.finder.FindImportedProtofilesInFile(p)
		if err != nil {
			return fmt.Errorf("unable to find dependencies in file %s:%v", p, err)
		}

		v.dependencies.EnqueueDependencies(importedFiles...)
	}

	for !v.dependencies.IsEmpty() {
		dependency := v.dependencies.Dequeue()

		for _, rule := range v.Rules {
			if strings.HasPrefix(dependency, rule.Prefix) {
				err := v.gitrepo.Clone(rule.RepoSSHCloneURL(), v.pathFormatter.FormatCachePath(rule.RepoDir()))
				if err != nil {
					return fmt.Errorf("unable to clone: %v", err)
				}

				repoProtoPath := v.pathFormatter.FormatCachePath(rule.DependencyDir(), dependency)
				importedFilename := v.pathFormatter.FormatOutputPath(dependency)
				importDir := filepath.Dir(importedFilename)

				if err := v.fs.CreateDirectory(importDir); err != nil {
					return fmt.Errorf("unable to create a directory for vendored file %s:%v", importDir, err)
				}

				if err := v.fs.CopyFile(repoProtoPath, importedFilename); err != nil {
					return fmt.Errorf("unable to copy file %s:%v", repoProtoPath, err)
				}

				dependencyImportedFiles, err := v.finder.FindImportedProtofilesInFile(repoProtoPath)
				if err != nil {
					return fmt.Errorf("unable to find dependencies in %s: %v", repoProtoPath, err)
				}

				v.dependencies.EnqueueDependencies(dependencyImportedFiles...)
			}
		}
	}

	return nil
}

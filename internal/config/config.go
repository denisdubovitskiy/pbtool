package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Project struct {
	ProtoDirs       []string             `yaml:"local_proto_dirs"`
	Build           []ExternalDependency `yaml:"build"`
	ExternalDepsDir string               `yaml:"external_dir"`
	ReplaceRules    []Rule               `yaml:"rules"`
	ProtoVendorDir  string               `yaml:"output_dir"`
	ExternalDeps    []ExternalDependency `yaml:"external"`
}

type Plugin struct {
	Name    string   `yaml:"name"`
	Path    string   `yaml:"path"`
	Options []string `yaml:"options"`
}

type ExternalDependency struct {
	RepoPath string   `yaml:"repo"`
	FilePath string   `yaml:"file_path"`
	Host     string   `yaml:"host"`
	Plugins  []Plugin `yaml:"plugins"`
}

func (e *ExternalDependency) RepoSSHCloneURL() string {
	return fmt.Sprintf("git@%s:%s.git", e.Host, e.RepoPath)
}

func (e *ExternalDependency) RepoDir() string {
	return filepath.Join(e.Host, e.RepoPath)
}

func (e *ExternalDependency) Filename() string {
	return filepath.Join(e.Host, e.RepoPath, e.FilePath)
}

type Rule struct {
	Prefix  string `yaml:"prefix"`
	Repo    string `yaml:"repo"`
	Subpath string `yaml:"subpath"`
	Host    string `yaml:"host"`
}

func (r *Rule) RepoDir() string {
	return filepath.Join(r.Host, r.Repo)
}

func (r *Rule) DependencyDir() string {
	return filepath.Join(r.RepoDir(), r.Subpath)
}

func (r *Rule) RepoSSHCloneURL() string {
	return fmt.Sprintf("git@%s:%s.git", r.Host, r.Repo)
}

func Read(path string) (*Project, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("config: unable to read config file: %s: %v", path, err)
	}

	conf := &Project{}

	if err := yaml.Unmarshal(b, conf); err != nil {
		return nil, fmt.Errorf("config: unable to unmarshal config file: %s: %v", path, err)
	}

	return conf, nil
}

func ReadAll(configFileName, repoCacheDir string) (*Project, error) {
	conf, err := Read(configFileName)
	if err != nil {
		return nil, fmt.Errorf("unable to read conf: %v", err)
	}

	if err := os.MkdirAll(conf.ProtoVendorDir, 0755); err != nil {
		return nil, fmt.Errorf("unable to create output dir %s: %v", conf.ProtoVendorDir, err)
	}

	if len(conf.ExternalDeps) > 0 {
		if err := os.MkdirAll(conf.ExternalDepsDir, 0755); err != nil {
			return nil, fmt.Errorf("unable to create external dir %s: %v", conf.ProtoVendorDir, err)
		}
	}

	if err := os.MkdirAll(repoCacheDir, 0755); err != nil {
		return nil, fmt.Errorf("unable to create a repo cache dir %s:%v", repoCacheDir, err)
	}

	return conf, nil
}

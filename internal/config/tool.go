package config

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type Tool struct {
	RepoCacheDir string
	ConfFile     string
}

func ReadToolConfig() (*Tool, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("toolconfig: unable to get a home home: %v", err)
	}

	conf := &Tool{}

	flag.StringVar(
		&conf.RepoCacheDir,
		"cachedir",
		filepath.Join(home, "proto-cache"),
		"A directory, which is used to cache dependency repositories",
	)

	flag.StringVar(
		&conf.ConfFile,
		"conf",
		"./vendor.yaml",
		"A config file which is used to configure vendoring",
	)

	flag.Parse()

	return conf, nil
}

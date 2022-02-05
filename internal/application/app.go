package application

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"

	"github.com/denisdubovitskiy/pbtool/internal/application/internal/command/build"
	"github.com/denisdubovitskiy/pbtool/internal/application/internal/command/setup"
	"github.com/denisdubovitskiy/pbtool/internal/application/internal/command/vendor"
)

type Runner interface {
	Run(arguments []string) (err error)
}

func New() (Runner, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("toolconfig: unable to get a home home: %v", err)
	}

	app := &cli.App{
		Name: "pbtool - A tool for working with protobuf dependencies",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "cachedir",
				Aliases: []string{"c"},
				Usage:   "Cache directory",
				Value:   filepath.Join(home, "proto-cache"),
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"f"},
				Usage:   "Project file",
				Value:   "./.pbtool.yaml",
			},
		},
		Commands: []*cli.Command{
			build.Command(),
			vendor.Command(),
			setup.Command(),
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app, nil
}

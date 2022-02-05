package vendor

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/denisdubovitskiy/pbtool/internal/application/internal/runner"
	"github.com/denisdubovitskiy/pbtool/internal/config"
	"github.com/denisdubovitskiy/pbtool/internal/vendorer"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:   "vendor",
		Usage:  "Vendor",
		Flags:  []cli.Flag{},
		Action: runner.RunWithConfig(vendor),
	}
}

func vendor(conf *config.Project, toolConf *config.Tool) error {
	err := vendorer.
		New(vendorer.OptsFromConfig(conf, toolConf.RepoCacheDir)).
		VendorDependencies()

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}

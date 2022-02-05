package setup

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/denisdubovitskiy/pbtool/internal/application/internal/appconf"
	"github.com/denisdubovitskiy/pbtool/internal/config/template"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Init",
		Action: func(context *cli.Context) error {
			filename := appconf.ConfigFileFromContext(context)

			if _, err := os.Stat(filename); err == nil {
				return nil
			}

			f, err := os.Create(filename)
			if err != nil {
				return err
			}

			defer f.Close()

			if _, err := f.WriteString(template.ConfTemplate); err != nil {
				return err
			}

			return nil
		},
	}
}

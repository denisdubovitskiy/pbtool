package runner

import (
	"fmt"

	"github.com/denisdubovitskiy/pbtool/internal/application/internal/appconf"
	"github.com/denisdubovitskiy/pbtool/internal/config"
	"github.com/urfave/cli/v2"
)

func RunWithConfig(fn func(conf *config.Project, toolConf *config.Tool) error) cli.ActionFunc {
	return func(context *cli.Context) error {
		conf, err := config.ReadAll(
			appconf.ConfigFileFromContext(context),
			appconf.CacheDirFromContext(context),
		)
		if err != nil {
			return fmt.Errorf("unable to read tool config: %v", err)
		}

		protoDirs := conf.ProtoDirs
		if len(conf.ExternalDepsDir) > 0 {
			protoDirs = append(protoDirs, conf.ExternalDepsDir)
		}

		conf.ProtoDirs = protoDirs

		return fn(conf, &config.Tool{
			RepoCacheDir: appconf.CacheDirFromContext(context),
			ConfFile:     appconf.ConfigFileFromContext(context),
		})
	}
}

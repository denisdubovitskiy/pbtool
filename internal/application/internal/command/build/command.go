package build

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/denisdubovitskiy/pbtool/internal/application/internal/runner"
	"github.com/denisdubovitskiy/pbtool/internal/config"
	"github.com/denisdubovitskiy/pbtool/internal/pathmanager"
	"github.com/denisdubovitskiy/pbtool/internal/protoc"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:   "build",
		Usage:  "Build",
		Flags:  []cli.Flag{},
		Action: runner.RunWithConfig(build),
	}
}

func build(conf *config.Project, toolConf *config.Tool) error {
	pathformatter := pathmanager.New(
		toolConf.RepoCacheDir,
		conf.ProtoVendorDir,
		conf.ExternalDepsDir,
	)

	for _, es := range conf.ExternalDeps {
		filePath := pathformatter.FormatExternalPath(es.Filename())
		if err := protoc.Compile(filePath, conf.ProtoVendorDir, encodePlugins(es.Plugins)); err != nil {
			return fmt.Errorf("unable to compile: %v", err)
		}
	}

	for _, es := range conf.Build {
		if err := protoc.Compile(es.FilePath, conf.ProtoVendorDir, encodePlugins(es.Plugins)); err != nil {
			return fmt.Errorf("unable to compile: %v", err)
		}
	}

	return nil
}

func encodePlugins(in []config.Plugin) []protoc.Plugin {
	plugins := make([]protoc.Plugin, len(in))

	for i, p := range in {
		plugins[i] = protoc.Plugin(p)
	}

	return plugins
}

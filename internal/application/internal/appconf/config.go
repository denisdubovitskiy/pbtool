package appconf

import "github.com/urfave/cli/v2"

func CacheDirFromContext(ctx *cli.Context) string {
	return ctx.String("cachedir")
}

func ConfigFileFromContext(ctx *cli.Context) string {
	return ctx.String("config")
}

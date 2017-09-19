package flags

import (
	"github.com/urfave/cli"
)

// See https://github.com/graymeta/stow/blob/master/google/config.go
var Google = []cli.Flag{
	cli.StringFlag{
		Name:   "gcs-json",
		Usage:  "GCS Service account JSON token",
		EnvVar: "PLUGIN_GCS_JSON",
	},
	cli.StringFlag{
		Name:   "gcs-project-id",
		Usage:  "GCS project id",
		EnvVar: "PLUGIN_GCS_PROJECT_ID",
	},
	cli.StringFlag{
		Name:   "gcs-scopes",
		Usage:  "GCS scopes",
		EnvVar: "PLUGIN_GCS_SCOPES",
	},
}

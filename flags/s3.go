package flags

import (
	"github.com/urfave/cli"
)

// See https://github.com/graymeta/stow/blob/master/s3/config.go
var S3 = []cli.Flag{
	cli.StringFlag{
		Name:   "S3-auth-type",
		Usage:  "S3 auth type",
		EnvVar: "PLUGIN_S3_AUTH_TYPE",
	},
	cli.StringFlag{
		Name:   "aws-access-key-id",
		Usage:  "AWS access key id",
		EnvVar: "PLUGIN_AWS_ACCESS_KEY_ID,CACHE_AWS_ACCESS_KEY_ID",
	},
	cli.StringFlag{
		Name:   "aws-secret-key",
		Usage:  "aws secret key",
		EnvVar: "PLUGIN_AWS_SECRET_KEY,CACHE_AWS_SECRET_KEY",
	},
	cli.StringFlag{
		Name:   "aws-region",
		Usage:  "aws region",
		EnvVar: "PLUGIN_AWS_REGION",
	},
	cli.StringFlag{
		Name:   "s3-endpoint",
		Usage:  "s3 endpoint",
		Value:  "s3.amazonaws.com",
		EnvVar: "PLUGIN_S3_ENDPOINT",
	},
	cli.StringFlag{
		Name:   "s3-disable-ssl",
		Usage:  "s3 disable ssl",
		EnvVar: "PLUGIN_S3_DISABLE_SSL",
	},
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/fixate/drone-stow-cache/flags"
	"github.com/fixate/drone-stow-cache/storage"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "cloud storage cache plugin"
	app.Usage = "cloud storage cache plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		// Plugin flags
		cli.StringSliceFlag{
			Name:   "mount",
			Usage:  "directories to cache",
			EnvVar: "PLUGIN_MOUNT",
		},
		cli.StringFlag{
			Name:   "kind",
			Usage:  "kind of storage provider supported by stow (e.g s3, google, azure etc)",
			EnvVar: "PLUGIN_KIND",
		},
		cli.StringFlag{
			Name:   "action",
			Usage:  "action to perform on the cache (restore, rebuild)",
			EnvVar: "PLUGIN_ACTION",
		},
		cli.BoolFlag{
			Name:   "flush",
			Usage:  "whether to flush the cache before rebuilding - only works when using the rebuild action",
			EnvVar: "PLUGIN_FLUSH",
		},
		cli.StringFlag{
			Name:   "flush-age",
			Usage:  "flush cache files older then # days",
			EnvVar: "PLUGIN_FLUSH_AGE",
			Value:  "30",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug plugin output",
			EnvVar: "PLUGIN_DEBUG",
		},

		// Cache location flags
		cli.StringFlag{
			Name:   "filename",
			Usage:  "filename for the cache",
			EnvVar: "PLUGIN_FILENAME",
		},
		cli.StringFlag{
			Name:   "prefix-path",
			Usage:  "appended onto the storage path",
			Value:  "/",
			EnvVar: "PLUGIN_PATH",
		},
		cli.StringFlag{
			Name:   "fallback-path",
			Usage:  "fallback path is used if the override path or implicit path does not exist",
			EnvVar: "PLUGIN_FALLBACK_PATH",
		},
		cli.StringFlag{
			Name:   "override-path",
			Usage:  "explicitly set the path to rebuild or restore",
			EnvVar: "PLUGIN_OVERRIDE_PATH",
		},

		// Stow flags (your storage location (url) and container)
		cli.StringFlag{
			Name:   "stow.location",
			Usage:  "stow location",
			EnvVar: "PLUGIN_STOW_LOCATION",
		},
		cli.StringFlag{
			Name:   "stow.container",
			Usage:  "container for storage e.g. bucket name",
			EnvVar: "PLUGIN_STOW_CONTAINER,DRONE_REPO_NAME",
		},

		// Repo flags
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repo owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repo name",
			EnvVar: "DRONE_REPO_NAME",
		},
	}

	// Provider flags
	app.Flags = append(app.Flags, flags.S3...)
	app.Flags = append(app.Flags, flags.Google...)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	// Set the logging level
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	kind := c.String("kind")
	if len(kind) == 0 {
		return errors.New("No kind specified - specify the kind of provider to use e.g s3, google etc")
	}

	// Determine the action for the plugin
	action := c.String("action")
	shouldFlush := c.Bool("flush")

	if len(action) == 0 {
		return errors.New("No action specified")
	}

	// Look for the mount points to rebuild
	mount := c.StringSlice("mount")

	if len(mount) == 0 {
		return errors.New("No mounts specified")
	}

	// Get the filename
	filename := c.GlobalString("filename")
	if len(filename) == 0 {
		log.Info("No filename specified. Using default 'cache.tar'")
		filename = "cache.tar"
	}

	path := getPath(c)
	fallbackPath := c.GlobalString("fallback-path")

	flushAge, err := strconv.Atoi(c.String("flush_age"))
	if err != nil {
		return err
	}

	config, err := storage.ContextToStowConfig(kind, c)
	if err != nil {
		return err
	}

	storage, err := storage.New(kind, config)

	p := &Plugin{
		Action:       action,
		Filename:     filename,
		Path:         path,
		FallbackPath: fallbackPath,
		FlushPath:    path,
		ShouldFlush:  shouldFlush,
		FlushAge:     flushAge,
		Mount:        mount,
		Storage:      storage,
	}

	return p.Exec()
}

func getPath(c *cli.Context) string {
	if path := c.GlobalString("override-path"); len(path) > 0 {
		return path
	}

	pathPrefix := c.String("path-prefix")

	return fmt.Sprintf(
		"/%s/%s/%s/%s",
		strings.Trim(pathPrefix, "/"),
		c.String("repo.owner"),
		c.String("repo.name"),
		c.String("commit.branch"),
	)
}

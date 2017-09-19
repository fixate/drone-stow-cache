package storage

import (
	"fmt"
	"io"

	"github.com/graymeta/stow"
	//"github.com/graymeta/stow/azure"
	"github.com/graymeta/stow/google"
	//"github.com/graymeta/stow/local"
	//"github.com/graymeta/stow/oracle"
	"github.com/graymeta/stow/s3"
	//"github.com/graymeta/stow/swift"
	"github.com/drone/drone-cache-lib/storage"
	"github.com/urfave/cli"
)

type Storage struct {
	Location *stow.Location
}

func New(kind string, config stow.Config) (storage.Storage, error) {
	location, err := stow.Dial(kind, config)
	return Storage{&location}, err
}

func (s Storage) Get(p string, dest io.Writer) error {
	return nil
}

func (s Storage) Put(p string, src io.Reader) error {
	return nil
}

func (s Storage) List(p string) ([]storage.FileEntry, error) {
	return []storage.FileEntry{}, nil
}

func (s Storage) Delete(p string) error {
	return nil
}

type Config stow.Config

func ContextToStowConfig(kind string, c *cli.Context) (config stow.ConfigMap, err error) {
	switch kind {
	case s3.Kind:
		config = stow.ConfigMap{}
	case google.Kind:
		config = stow.ConfigMap{
			google.ConfigJSON:      c.String("gcs-json"),
			google.ConfigProjectId: c.String("gcs-project-id"),
			google.ConfigScopes:    c.String("gcs-scopes"),
		}
	default:
		err = fmt.Errorf("Unsupported kind '%s'", kind)
	}
	return
}

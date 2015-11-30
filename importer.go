package main

import (
	"errors"

	"github.com/codegangsta/cli"
)

var IMPORTERS ImporterConstructors = ImporterConstructors{
	"redis": ConstructRedisImporter,
}

type ImporterConstructors map[string]func(*cli.Context) Importer

type Importer interface {
	Import(map[string]interface{}) error
}

func ConstructImporter(c *cli.Context, importers ImporterConstructors) (Importer, error) {
	if v, found := importers[c.String("backend")]; found {
		return v(c), nil
	} else {
		return nil, errors.New("Backend not supported")
	}
}

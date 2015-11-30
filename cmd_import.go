package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var cmdImport cli.Command = cli.Command{
	Name: "import",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "prefix",
			Value: "/",
			Usage: "key path prefix",
		},
		cli.StringFlag{
			Name:  "json",
			Value: "",
			Usage: "json file to import. (if empty stdin will be used)",
		},
		cli.StringFlag{
			Name:  "backend",
			Value: "redis",
			Usage: "name of the confd backend. Ex: redis, consul, etcd",
		},
		cli.StringFlag{
			Name:  "node",
			Value: "",
			Usage: "address and port of the backend. ex: 127.0.0.1:6379",
		},
	},
	Usage: "",
	Action: func(c *cli.Context) {
		importer, err := ConstructImporter(c, IMPORTERS)
		if err != nil {
			fmt.Println(err)
			return
		}

		importer.Import(map[string]interface{}{
			"name": "myname",
		})
	},
}

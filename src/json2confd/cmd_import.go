package main

import (
	"fmt"
	"io"
	"log"
	"os"

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
			Name:  "file",
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
		cli.StringFlag{
			Name:  "keys",
			Value: "false",
			Usage: "print the imported keys from the source file (possible values: true/false)",
		},
	},
	Usage: "Import the json into the specified backend",
	Action: func(c *cli.Context) {
		var err error
		var keys []string
		var importer Importer
		var reader io.Reader
		var flat map[string]interface{}

		if importer, err = ConstructImporter(c, IMPORTERS); err != nil {
			log.Fatal(err)
		}

		if reader, err = getJsonReader(c.String("file")); err != nil {
			log.Fatal(err)
		}

		if flat, err = FlattenReader(reader, "/"); err != nil {
			log.Fatal(err)
		}

		if err, keys = importer.Import(flat); err != nil {
			log.Fatal(err)
		}

		if c.String("keys") == "true" {
			fmt.Println(beautifyKeys(keys))
		}

		fmt.Printf("Success! Imported %v entries into %s(%s)\n", len(flat), c.String("backend"), c.String("node"))
	},
}

func getJsonReader(file string) (io.Reader, error) {
	if file == "" {
		return os.Stdin, nil
	} else {
		return os.Open(file)
	}
}

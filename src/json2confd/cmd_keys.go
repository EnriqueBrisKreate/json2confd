package main

import (
	"github.com/codegangsta/cli"
	"io"
	"log"
	"os"
)

var cmdKeys cli.Command = cli.Command{
	Name: "keys",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "file",
			Value: "",
			Usage: "json file to import. (if empty stdin will be used)",
		},
		cli.StringFlag{
			Name:  "output",
			Value: "json",
			Usage: "output format: json",
		},
	},
	Usage: "Output the keys to be imported",
	Action: func(c *cli.Context) {
		var (
			reader   io.Reader
			output   outputData
			outputer outputer
			err      error
		)

		if reader, err = os.Open(c.String("file")); err != nil {
			log.Fatal(err)
		}
		if output, err = extractKeys(reader); err != nil {
			log.Fatal(err)
		}
		if outputer, err = getOutputer(c.String("output")); err != nil {
			log.Fatal(err)
		}

		outputer.output(output, c.App.Writer)

	},
}

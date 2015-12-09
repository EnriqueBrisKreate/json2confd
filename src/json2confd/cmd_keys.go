package main

import (
	"github.com/codegangsta/cli"
	"io"
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

		var reader io.Reader
		var err error

		if reader, err = os.Open(c.String("file")); err != nil {
			// TODO :: outputs the error to the io.Writer
		} else {

			if file_map, err := jsonToMap(reader); err != nil {
				// TODO :: outputs the error to the io.Writer
			} else {
				keys := make([]string, len(file_map))
				i := 0
				for k := range file_map {
					keys[i] = k
					i = i + 1
				}

				output := OutputData{
					Keys: keys,
				}

				if err := Output(output, c.String("output")); err != nil {
					// TODO :: outputs the error to the io.Writer

				}
			}
		}

	},
}
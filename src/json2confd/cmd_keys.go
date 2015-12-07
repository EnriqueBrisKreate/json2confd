package main

import (
	"github.com/codegangsta/cli"
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

		if file_map, err := jsonToMap(c.String("file")); err != nil {
			panic(err)
		} else {
			le := len(file_map)
			keys := make([]string, 0, le)
			for k := range file_map {
				keys = append(keys, k)
			}

			output := OutputData{
				Keys: keys,
			}

			Output(output, c.String("output"))
		}

	},
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

var cmdImport cli.Command = cli.Command{
	Name: "import",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "onetime",
			Usage: "Run once or keeps watching the json file for changes.",
		},
		cli.IntFlag{
			Name:  "interval",
			Value: 60,
			Usage: "Number of seconds in which the file will be checked for changes.",
		},
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
	},
	Usage: "Import the json into the specified backend",
	Action: func(c *cli.Context) {
		interval, err := time.ParseDuration(fmt.Sprint(c.Int("interval")) + "s")
		if err != nil {
			log.Fatal(err)
		}

		ip := ImportParams{
			onetime:  c.Bool("onetime"),
			interval: interval,
			prefix:   c.String("prefix"),
			file:     c.String("file"),
			backend:  c.String("backend"),
			node:     c.String("node"),
		}
		if ip.onetime == true {
			err = BasicImporter{}.Import(ip)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = ContinuousImporter{}.Import(ip)
			if err != nil {
				log.Println(err)
			}
		}

		//fmt.Printf("Success! Imported %v entries into %s(%s)\n", len(flat), ip.backend, ip.node)
	},
}

func getJsonReader(file string) (io.Reader, error) {
	if file == "" {
		return os.Stdin, nil
	} else {
		return os.Open(file)
	}
}

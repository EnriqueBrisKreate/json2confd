package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "json2consul"
	app.Usage = ""
	app.Commands = []cli.Command{
		cmdImport,
		cmdKeys,
	}

	app.Run(os.Args)
}

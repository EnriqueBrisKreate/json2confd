package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func ConstructRedisImporter(c *cli.Context) Importer {
	return RedisImporter{}
}

type RedisImporter struct {
}

func (r RedisImporter) Import(p map[string]interface{}) error {
	fmt.Println("Importing to redis")
	fmt.Printf("%#v", p)
	return nil
}

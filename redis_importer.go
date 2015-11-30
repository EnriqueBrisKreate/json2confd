package main

import (
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/garyburd/redigo/redis"
	"strings"
)

func ConstructRedisImporter(c *cli.Context) Importer {
	server := c.String("node")
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return RedisImporter{
		pool:   pool,
		prefix: c.String("prefix"),
	}
}

type RedisImporter struct {
	pool   *redis.Pool
	prefix string
}

func (r RedisImporter) Import(p map[string]interface{}) error {

	prefix := ""
	if len(r.prefix) > 0 && r.prefix != "/" {
		prefix = "/" + strings.TrimSpace(strings.Trim(r.prefix, "/"))
	}

	for k, v := range p {

		_, err := r.pool.Get().Do("SET", prefix + k, fmt.Sprint(v))
		if err != nil {
			return err
		}

	}
	
	return nil
}

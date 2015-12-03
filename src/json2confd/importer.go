package main

import (
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type ImportParams struct {
	onetime  bool
	interval time.Duration
	prefix   string
	file     string
	backend  string
	node     string
}

type Importer interface {
	Import(ImportParams) error
}

type BasicImporter struct {
}

func (b BasicImporter) Import(ip ImportParams) error {
	var err error
	var backend Backend
	var reader io.Reader
	var flat map[string]interface{}

	if backend, err = ConstructBackend(ip, BACKENDS); err != nil {
		return err
	}

	if reader, err = getJsonReader(ip.file); err != nil {
		return err
	}

	if flat, err = FlattenReader(reader, "/"); err != nil {
		return err
	}

	if err = backend.Import(flat); err != nil {
		return err
	}
	return nil
}

type ContinuousImporter struct {
}

func (c ContinuousImporter) Import(ip ImportParams) error {
	var err error
	var backend Backend
	var reader io.Reader
	var flat map[string]interface{}

	if backend, err = ConstructBackend(ip, BACKENDS); err != nil {
		return err
	}

	tmp := ""
	for {
		if reader, err = os.Open(ip.file); err != nil {
			log.Println(err)
		}
		bytes, err := ioutil.ReadAll(reader)

		encoded := base64.StdEncoding.EncodeToString(bytes)
		if tmp != encoded {
			if flat, err = FlattenJson(bytes, "/"); err != nil {
				log.Println(err)
			}

			if err = backend.Import(flat); err != nil {
				log.Println(err)
			}
			tmp = encoded
			log.Println("Imported", flat)
		}
		time.Sleep(ip.interval)
	}
	return nil
}

var BACKENDS BackendConstructors = BackendConstructors{
	"redis": ConstructRedisImporter,
}

var BACKEND_PORTS map[string]string = map[string]string{
	"redis": "6379",
}

type BackendConstructors map[string]func(ImportParams) Backend

type Backend interface {
	Import(map[string]interface{}) error
}

func ConstructBackend(ip ImportParams, backends BackendConstructors) (Backend, error) {
	if v, found := backends[ip.backend]; found {
		return v(ip), nil
	} else {
		return nil, errors.New("Backend not supported")
	}
}

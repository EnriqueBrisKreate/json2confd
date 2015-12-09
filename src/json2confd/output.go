package main

import (
	"encoding/json"
	"fmt"
)

type OutputData struct {
	Keys []string `json:"keys"`
}

func Output(data OutputData, outputFormat string) (err error) {
	switch outputFormat {
	case "json":
		err = outputJson(data)
	}

	return
}

func outputJson(data OutputData) error {
	if retJson, err := json.Marshal(data); err != nil {
		return err
	} else {
		// TODO :: outputs to the io.Writer
		fmt.Printf("%s\n", retJson)
	}

	return nil
}
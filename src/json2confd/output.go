package main

import (
	"encoding/json"
	"fmt"
)

type OutputData struct {
	Keys []string `json:"keys"`
}

func Output(data OutputData, outputFormat string) {
	switch outputFormat {
	case "json":
		outputJson(data)
	}
}

func outputJson(data OutputData) {
	if retJson, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		fmt.Printf("%s\n", retJson)
	}
}

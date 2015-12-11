package main

import (
	"errors"
	"io"
)

type outputData struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Keys    []string `json:"keys"`
}

type outputer interface {
	output(data outputData, writer io.Writer) error
}

func getOutputer(format string) (outputer, error) {
	switch format {
	case "json":
		return jsonOutputer{}, nil
	default:
		return nil, errors.New("Format not recognized")
	}
}

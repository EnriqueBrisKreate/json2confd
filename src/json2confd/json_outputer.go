package main

import (
	"encoding/json"
	"io"
)

type jsonOutputer struct{}

func (j jsonOutputer) output(data outputData, writer io.Writer) error {
	if retJson, err := json.Marshal(data); err != nil {
		return err
	} else {
		_, err:= writer.Write(retJson)
		return err
	}
}

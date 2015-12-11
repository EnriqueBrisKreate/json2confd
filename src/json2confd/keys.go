package main

import "io"

func extractKeys(reader io.Reader) (outputData, error) {
	if file_map, err := jsonToMap(reader); err != nil {
		return outputData{Success: false}, err
	} else {
		output := outputData{
			Success: true,
			Keys:    make([]string, len(file_map)),
		}
		i := 0
		for k := range file_map {
			output.Keys[i] = k
			i = i + 1
		}

		return output, nil
	}
}

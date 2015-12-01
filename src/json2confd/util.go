package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func FlattenReader(reader io.Reader, d string) (map[string]interface{}, error) {
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return FlattenJson(contents, d)
}

// FlattenJson takes in a  nested JSON as a byte array and a delimiter and returns an
// exploded/flattened json byte array
func FlattenJson(b []byte, d string) (map[string]interface{}, error) {
	var input map[string]interface{}
	var err error
	err = json.Unmarshal(b, &input)
	if err != nil {
		return nil, err
	}
	return explodeMap(input, "", d)
}

// FlattenJsonStr explodes a nested JSON string to an unnested one
// parameters to pass to the function are
// * s: the JSON string
// * d: the delimiter to use when unnesting the JSON object.
// {"person":{"name":"Joe", "address":{"street":"123 Main St."}}}
// explodes to:
// {"person.name":"Joe", "person.address.street":"123 Main St."}
func FlattenJsonStr(s string, d string) (map[string]interface{}, error) {
	b := []byte(s)
	out, err := FlattenJson(b, d)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func explodeList(l []interface{}, parent string, delimiter string) (map[string]interface{}, error) {
	var err error
	var key string
	j := make(map[string]interface{})
	for k, i := range l {
		if len(parent) > 0 {
			key = parent + delimiter + strconv.Itoa(k)
		} else {
			key = strconv.Itoa(k)
		}
		switch v := i.(type) {
		case nil:
			j[key] = v
		case int:
			j[key] = v
		case float64:
			j[key] = v
		case string:
			j[key] = v
		case bool:
			j[key] = v
		case []interface{}:
			out := make(map[string]interface{})
			out, err = explodeList(v, key, delimiter)
			if err != nil {
				return nil, err
			}
			for newkey, value := range out {
				j[newkey] = value
			}
		case map[string]interface{}:
			out := make(map[string]interface{})
			out, err = explodeMap(v, key, delimiter)
			if err != nil {
				return nil, err
			}
			for newkey, value := range out {
				j[newkey] = value
			}
		default:
			// do nothing
		}
	}
	return j, nil
}

func explodeMap(m map[string]interface{}, parent string, delimiter string) (map[string]interface{}, error) {
	var err error
	j := make(map[string]interface{})
	for k, i := range m {
		if len(parent) > 0 {
			k = parent + delimiter + k
		} else {
			k = delimiter + k
		}
		switch v := i.(type) {
		case nil:
			j[k] = v
		case int:
			j[k] = v
		case float64:
			j[k] = v
		case string:
			j[k] = v
		case bool:
			j[k] = v
		case []interface{}:

			tmp, err := json.Marshal(v)

			if err != nil {
				return nil, err
			}

			j[k] = fmt.Sprintf("%s", tmp)

		case map[string]interface{}:
			out := make(map[string]interface{})
			out, err = explodeMap(v, k, delimiter)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				j[key] = value
			}
		default:
			//nothing
		}
	}
	return j, nil
}

func ensurePort(address string, backend string) string {
	if address != "" && strings.Contains(address, ":") == false {
		if port, found := IMPORTERS_PORTS[backend]; found {
			return address + ":" + port
		}
	}
	return address
}

package main

import (
	. "gopkg.in/check.v1"
	"strings"
)

var _ = Suite(&TestKeysSuite{})

type TestKeysSuite struct {
	json1 string
}

func (s *TestKeysSuite) SetUpTest(c *C) {
	s.json1 = `
	{
		"description": "a schema given for items",
		"schema": {
			"items": {"type": "integer"}
		},
		"test1": {
			"description": "valid items",
			"data": [ 1, 2, 3 ],
			"valid": true
		},
		"test3": {
			"description": "ignores non-arrays",
			"data": {"foo" : "bar"}
		}
	}
	`
}

// Test calling extractKeys with a valid json
func (s *TestKeysSuite) TestExtractKeys_ValidJson(c *C) {

	expected := []string{
		"/test1/valid",
		"/test3/description",
		"/test3/data/foo",
		"/description",
		"/schema/items/type",
		"/test1/description",
		"/test1/data",
	}

	reader := strings.NewReader(s.json1)
	output, err := extractKeys(reader)

	// no error
	c.Assert(err, IsNil)
	// successfully operation
	c.Assert(output.Success, Equals, true)
	// right number of keys
	c.Assert(len(output.Keys), Equals, len(expected))

	// TODO ::: Make a fn to compare two []string with unsorted values (put the following code into a fn)
	// load the strings into a map (O(n)) to later compare (O(1) each compare operation)
	stringMap := make(map[string]bool, len(output.Keys))
	for _, v := range output.Keys {
		stringMap[v] = true
	}
	// check the keys
	contains_all_keys := true
	for _, v := range expected {
		if !stringMap[v] {
			contains_all_keys = false
			break
		}
	}
	c.Assert(contains_all_keys, Equals, true)

}

// Test calling extractKeys with an invalid json
func (s *TestKeysSuite) TestExtractKeys_InvalidJson(c *C) {
	reader := strings.NewReader("abcde")
	output, err := extractKeys(reader)

	c.Assert(err, NotNil)
	c.Assert(output.Success, Equals, false)
}

// Test calling extractKeys with an empty json
func (s *TestKeysSuite) TestExtractKeys_EmptyJson(c *C) {
	reader := strings.NewReader("")
	output, err := extractKeys(reader)

	c.Assert(err, NotNil)
	c.Assert(output.Success, Equals, false)
}
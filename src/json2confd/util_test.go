package main

import (
	"encoding/json"
	"fmt"
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&TestUtilSuite{})

type TestUtilSuite struct {
	json1            string
	map1             map[string]interface{}
	json2            string
	map2             map[string]interface{}
	json2ExpectedMap map[string]interface{}
}

func (s *TestUtilSuite) SetUpTest(c *C) {
	s.json1 = `
	{
		"name": "myname",
		"age": 23,
		"like_chocolate": false
	}
	`
	json.Unmarshal([]byte(s.json1), &s.map1)

	s.json2 = `
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
		"test2": {
			"description": "wrong type of items",
			"data": [1, "x"],
			"valid": false
		},
		"test3": {
			"description": "ignores non-arrays",
			"data": {"foo" : "bar"},
			"valid": true
		}
	}
	`
	json.Unmarshal([]byte(s.json2), &s.map2)

	s.json2ExpectedMap = map[string]interface{}{
		"/description":       "a schema given for items",
		"/schema/items/type": "integer",
		"/test1/description": "valid items",
		"/test1/data":        "[1,2,3]",
		"/test1/valid":       "true",
		"/test2/description": "wrong type of items",
		"/test2/data":        "[1,\"x\"]",
		"/test2/valid":       "false",
		"/test3/description": "ignores non-arrays",
		"/test3/data/foo":    "bar",
		"/test3/valid":       "true",
	}
}

// Test calling FlattenJsonStr with an invalid json
func (s *TestUtilSuite) TestFlattenJsonStrInvalidJSON(c *C) {
	res, err := FlattenJsonStr("This is an invalid json", "/")
	c.Assert(res, IsNil)
	c.Assert(err, NotNil)
}

// Test calling explodeMap with s.json1
func (s *TestUtilSuite) TestExplodeMapJson1(c *C) {
	expected := map[string]interface{}{
		"/name":           "myname",
		"/age":            "23",
		"/like_chocolate": "false",
	}

	res, err := explodeMap(s.map1, "", "/")
	c.Assert(err, IsNil)
	for k, v := range expected {
		_, key_exists := res[k]
		c.Assert(key_exists, Equals, true, Commentf("key: %#v, not found on res", k))
		c.Assert(fmt.Sprint(res[k]), Equals, v, Commentf("for k: %#v, result should be: %#v", k, v))
	}
}

// Test calling explodeMap with s.json2
func (s *TestUtilSuite) TestExplodeMapJson2(c *C) {
	res, err := explodeMap(s.map2, "", "/")
	c.Assert(err, IsNil)
	for k, v := range s.json2ExpectedMap {
		_, key_exists := res[k]
		c.Assert(key_exists, Equals, true, Commentf("key: %#v, not found on res ... %#v", k, res))
		c.Assert(fmt.Sprint(res[k]), Equals, v, Commentf("for k: %#v, result should be: %#v ... %#v", k, v, res[k]))
	}
}

// Test calling jsonToMap with a valid json
func (s *TestUtilSuite) TestJsonToMap_ValidJson(c *C) {

	reader := strings.NewReader(s.json2)
	res, err := jsonToMap(reader)

	c.Assert(res, NotNil)
	c.Assert(err, IsNil)

	// compare maps
	eq := myDeepMapEqual(res, s.json2ExpectedMap)
	c.Assert(eq, Equals, true)
}

// Test calling jsonToMap with an invalid json
func (s *TestUtilSuite) TestJsonToMap_InvalidJson(c *C) {
	reader := strings.NewReader("abcde")
	res, err := jsonToMap(reader)

	c.Assert(err, NotNil)
	c.Assert(len(res), Equals, 0)
}

// Test calling jsonToMap with an empty json
func (s *TestUtilSuite) TestJsonToMap_EmptyBodyJson(c *C) {
	reader := strings.NewReader("")
	res, err := jsonToMap(reader)

	c.Assert(err, NotNil)
	c.Assert(len(res), Equals, 0)
}

// Test calling MyDeepMapEqual with two equivalent maps (two maps with the same key&value combinations but not in the same positions)
func (s *TestUtilSuite) TestMyDeepMapEqual_ValidAndEquivalentMaps(c *C) {
	json2ExpectedMapDifferentSorting := map[string]interface{}{
		"/test2/description": "wrong type of items",
		"/test3/data/foo":    "bar",
		"/test3/description": "ignores non-arrays",
		"/test1/description": "valid items",
		"/description":       "a schema given for items",
		"/schema/items/type": "integer",
		"/test2/valid":       "false",
		"/test1/data":        "[1,2,3]",
		"/test1/valid":       "true",
		"/test3/valid":       "true",
		"/test2/data":        "[1,\"x\"]",
	}

	res := myDeepMapEqual(s.json2ExpectedMap, json2ExpectedMapDifferentSorting)

	c.Assert(res, Equals, true)
}

// Test calling MyDeepMapEqual with two maps with different len
func (s *TestUtilSuite) TestMyDeepMapEqual_DifferentLenMaps(c *C) {
	different_map := map[string]interface{}{
		"any_key": []interface{}{"a", 3},
	}

	res := myDeepMapEqual(s.json2ExpectedMap, different_map)

	c.Assert(res, Equals, false)
}

// Test calling MyDeepMapEqual with two maps (same len but different key&value combinations)
func (s *TestUtilSuite) TestMyDeepMapEqual_SameLenButDifferentContent(c *C) {
	json2ExpectedMapDifferentSorting := map[string]interface{}{
		"/test2/description": "wrong type of items",
		"/test3/data/foo":    "bar",
		"/test3/description": "ignores non-arrays",
		"/test1/description": "valid items",
		"/description_new":   "a schema given for items",
		"/schema/items/type": 12345,
		"/test2/valid":       "false",
		"/test1/data":        "[1,2,3]",
		"/test1/valid":       true,
		"/test3/valid":       "true",
		"/test2/data":        "[1,\"x\"]",
	}

	res := myDeepMapEqual(s.json2ExpectedMap, json2ExpectedMapDifferentSorting)

	c.Assert(res, Equals, false)
}

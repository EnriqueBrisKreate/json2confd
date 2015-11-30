package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

import . "gopkg.in/check.v1"

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&TestUtilSuite{})

type TestUtilSuite struct {
	json1 string
	map1  map[string]interface{}
	json2 string
	map2  map[string]interface{}
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
}

// Test calling FlattenJsonStr with an invalid json
func (s *TestUtilSuite) TestFlattenJsonStrInvalidJSON(c *C) {
	res, err := FlattenJsonStr("This is an invalid json", "/")
	c.Assert(res, Equals, "")
	c.Assert(res, NotNil)
	fmt.Println(res, err)
}

// Test calling FlattenJsonStr with s.json1
func (s *TestUtilSuite) TestFlattenJsonStrJson1(c *C) {
	expected := map[string]interface{}{
		"/name":           "myname",
		"/age":            "23",
		"/like_chocolate": "false",
	}
	res, err := explodeMap(s.map1, "", "/")
	c.Assert(err, IsNil)
	for k, v := range expected {
		_, key_exists := res[k]
		c.Assert(key_exists, Equals, true)
		c.Assert(fmt.Sprint(res[k]), Equals, v, Commentf("for k: %#v, result should be: %#v", k, v))
	}
}

// Test calling FlattenJsonStr with s.json2
func (s *TestUtilSuite) TestFlattenJsonStrJson2(c *C) {
	expected := map[string]interface{}{
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
	res, err := explodeMap(s.map2, "", "/")
	c.Assert(err, IsNil)
	for k, v := range expected {
		_, key_exists := res[k]
		c.Assert(key_exists, Equals, true)
		c.Assert(fmt.Sprint(res[k]), Equals, v, Commentf("for k: %#v, result should be: %#v", k, v))
	}
}

package main

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&TestRedisImporterSuite{})

type TestRedisImporterSuite struct{}

func (s *TestRedisImporterSuite) SetUpTest(c *C) {

}

// Test calling FlattenJsonStr with an invalid json
func (s *TestRedisImporterSuite) TestFlattenJsonStrInvalidJSON(c *C) {

}

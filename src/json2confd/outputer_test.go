package main

import (
	"fmt"
	. "gopkg.in/check.v1"
)

var _ = Suite(&TestOutputSuite{})

type TestOutputSuite struct {
}

func (s *TestOutputSuite) SetUpTest(c *C) {
}

// Test calling getOutputer with a known string parameter (json)
func (s *TestUtilSuite) TestGetOutputer_KnownString(c *C) {
	outputer, err := getOutputer("json")

	c.Assert(err, IsNil)
	// check type
	c.Assert(fmt.Sprintf("%T", outputer), Equals, fmt.Sprintf("%T", jsonOutputer{}))
}

// Test calling getOutputer with an unknown string parameter (abcde)
func (s *TestUtilSuite) TestGetOutputer_UnknownString(c *C) {
	_, err := getOutputer("abcde")

	c.Assert(err, NotNil)
}

// Test calling getOutputer with an empty string parameter (abcde)
func (s *TestUtilSuite) TestGetOutputer_EmptyString(c *C) {
	_, err := getOutputer("")

	c.Assert(err, NotNil)
}

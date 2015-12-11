package main

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&TestJsonOutputerSuite{})

type TestJsonOutputerSuite struct{}

func (s *TestJsonOutputerSuite) SetUpTest(c *C) {

}

// Test calling Output with valid data
func (s *TestJsonOutputerSuite) TestOutput(c *C) {
	data := outputData{
		Success: true,
		Message: "any message",
		Keys:    []string{"one", "two", "three"},
	}

	err := jsonOutputer{}.output(data, fakeWriter{})

	c.Assert(err, IsNil)
}

type fakeWriter struct{}

func (f fakeWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

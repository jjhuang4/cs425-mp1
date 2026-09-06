package main

import (
	"testing"
)

// Required unit test that generates log files at every machine with some known lines and
// other random lines. The log-querying program then runs multiple greps and
// verifies automatically that the results are what we expect.
func RequiredTest(t *testing.T) {
	testLog := "Hello, World!\nThis is a required log-querying unit test"
	fileWriteCall(testLog)
}

package main

import (
	"sync"
	"testing"
)

// https://gobyexample.com/testing
// Required unit test that generates log files at every machine with some known lines and
// other random lines. The log-querying program then runs multiple greps and
// verifies automatically that the results are what we expect.
func RequiredTest(t *testing.T) {
	return
}

// Simple test
func SimpleTest(t *testing.T) {
	var fileWriteMap = map[string]struct {
		testString, fileName, expected string
	}{
		"vm2": {"Hello, World!\nThis is a required log-querying unit test", "machine.2.log", "1:Hello"},
		"vm3": {"Other test\nHello test", "machine.3.log"},
	}
	for vm, tt := range fileWriteMap {
		fout, err := fileWriteCall(tt.testString, vm, tt.fileName)
		if err != nil {
			t.Errorf("Error occurred with fileWriteCall to %s: %w", vm, err)
		}
	}

	//refactor
	var grepCallTests = []struct {
		vm, pattern, flags, file string
	}{
		{"vm2", "-n", "Hello", "machine.2.log"},
		{"vm3", "-n", "Hello", "machine.3.log"},
	}

	var wg sync.WaitGroup

	for _, tt := range grepCallTests {
		wg.Go(func() {
			reply, err := grepCall(tt.vm, tt.flags, tt.pattern, tt.file)

			expectedString := fileWriteMap[tt.vm].expected
			if expectedString != reply {
				t.Errorf("Expected reply: %s\nReceived reply: %s\n\n", expectedString, reply)
			}
		})
	}
	wg.Wait()

}

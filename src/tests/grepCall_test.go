package main

import (
	"cs425/mp1/src/client"
	"fmt"
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
func TestSimple(t *testing.T) {
	fmt.Println("HERE")
	var fileWriteMap = map[string]struct {
		testString, fileName, expected string
	}{
		"vm2": {"Hello, World!\nThis is a required log-querying unit test", "machine.2.log", "1:Hello, World!"},
		"vm3": {"Other test\nHello test", "machine.3.log", "2:Hello test"},
	}
	for vm, tt := range fileWriteMap {
		_, err := client.FileWriteCall(tt.testString, vm, tt.fileName)
		if err != nil {
			t.Errorf("\nError occurred with fileWriteCall to %s: %s", vm, err)
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
			reply, err := client.GrepCall(tt.vm, tt.flags, tt.pattern, tt.file)
			if err != nil {
				t.Errorf("Error when grepCall: %s", err)
			}
			expectedString := fileWriteMap[tt.vm].expected
			if expectedString+string('\n') != reply {
				fmt.Println(reply)
				t.Errorf("\nExpected reply: %s\nReceived reply: %s\n\n", expectedString, reply)
			}
		})
	}
	wg.Wait()

}

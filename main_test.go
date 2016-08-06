package main

import (
	"fmt"
	"testing"
)

func panicIfErr(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func TestCheckYaml(t *testing.T) {
	opts := &CmdOpts{}
	opts.Silent = true

	// Test valid yaml
	err := checkYaml([]byte("foo: bar\n"), opts)
	if err != nil {
		t.Errorf("Failed to check Valid yaml! Error = %v", err)
	}

	// Test invalid yaml
	err = checkYaml([]byte("foo\n"), opts)
	if err == nil {
		t.Error("Failed to check Invalid yaml! No error detected!")
	}
}

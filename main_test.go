package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	okYml := "tmp/ok.yml"
	okContent := []byte("foo: bar\n")
	err := ioutil.WriteFile(okYml, okContent, 0644)
	defer os.Remove(okYml)
	err = checkYaml([]string{okYml}, opts)
	if err != nil {
		t.Errorf("Failed to check okYml! Error = %v", err)
	}

	// Test invalid yaml
	ngYml := "tmp/ng.yml"
	ngContent := []byte("foo\n")
	err = ioutil.WriteFile(ngYml, ngContent, 0644)
	defer os.Remove(ngYml)
	err = checkYaml([]string{ngYml}, opts)
	if err == nil {
		t.Error("Failed to check ngYml! No error detected!")
	}
}

// Copyright (c) 2016 IKEDA Kiyoshi
// The MIT License (MIT)

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const version = "0.1.0"

const (
	exitOk = iota
	exitErr
)

type CmdOpts struct {
	Silent  bool `short:"s" long:"silent" description:"Don't print result"`
	Version bool `short:"v" long:"version" description:"Show version number"`
}

func main() {
	opts := &CmdOpts{}
	args, err := flags.ParseArgs(opts, os.Args[1:])
	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			os.Exit(exitErr)
		} else {
			croak(err, opts.Silent)
		}
	}

	if opts.Version {
		fmt.Printf("Version: %s\n", version)
		os.Exit(exitOk)
	}

	var buf []byte
	buf, err = getInputBuffer(args, opts)

	err = checkYaml(buf, opts)
	if err != nil {
		croak(err, opts.Silent)
	}

	os.Exit(exitOk)
}

func croak(e error, silent bool) {
	if !silent {
		fmt.Printf("%v\n", e)
	}
	os.Exit(exitErr)
}

func checkYaml(buf []byte, opts *CmdOpts) error {
	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal(buf, &m)
	if err != nil {
		return err
	}

	if !opts.Silent {
		fmt.Println("Syntax OK")
	}

	return nil
}

func getInputBuffer(args []string, opts *CmdOpts) ([]byte, error) {
	var (
		buf []byte
		err error
	)
	if len(args) > 0 {
		buf, err = readFile(args[0], opts.Silent)
	} else {
		buf = readStdin()
	}

	return buf, err
}

func readFile(path string, silent bool) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		croak(err, silent)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	return buf, err
}

func readStdin() []byte {
	var b bytes.Buffer
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		b.Write(sc.Bytes())
	}
	return b.Bytes()
}

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

const version = "0.9.0"

const (
	exitOk = iota
	exitErr
)

type CmdOpts struct {
	Silent  bool `short:"s" long:"silent" description:"Don't print result"`
	Version bool `short:"v" long:"version" description:"Show version number"`
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("ERROR!\n%v\n", err)
		}
		os.Exit(exitErr)
	}()

	opts := &CmdOpts{}
	args, err := flags.ParseArgs(opts, os.Args[1:])
	if err != nil {
		if _, ok := err.(*flags.Error); ok {
			os.Exit(exitErr)
		} else {
			panic(err)
		}
	}

	if opts.Version {
		fmt.Printf("Version: %s\n", version)
		os.Exit(exitOk)
	}

	buf := getInputBuffer(args, opts)

	err = checkYaml(buf, opts)
	croakIfError(err, opts.Silent)

	os.Exit(exitOk)
}

func croakIfError(e error, silent bool) {
	if e == nil {
		return
	}
	if silent {
		panic(nil)
	} else {
		panic(e)
	}
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

func getInputBuffer(args []string, opts *CmdOpts) []byte {
	var buf []byte
	if len(args) > 0 {
		buf = readFile(args[0], opts.Silent)
	} else {
		buf = readStdin()
	}

	return buf
}

func readFile(path string, silent bool) []byte {
	file, err := os.Open(path)
	croakIfError(err, silent)
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return buf
}

func readStdin() []byte {
	var b bytes.Buffer
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		b.Write(sc.Bytes())
	}
	return b.Bytes()
}

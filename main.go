package main

import (
	"bufio"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	var buf []byte
	if len(os.Args) > 1 {
		buf = readFile(os.Args[1])
	} else {
		buf = readStdin()
	}

	m := make(map[interface{}]interface{})
	err := yaml.Unmarshal(buf, &m)
	if err != nil {
		croak(err)
	}

	os.Exit(0)
}

func croak(e error) {
	fmt.Print(e)
	os.Exit(1)
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		croak(err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		croak(err)
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

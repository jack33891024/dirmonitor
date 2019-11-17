package config

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseConfig(t *testing.T) {
	fileName := "demo.cnf"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatalf("%v not found,err: %v\n", fileName, err)
	}

	line := strings.Split(string(data), "\r\n")

	fmt.Printf("%#v", line)
}

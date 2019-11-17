package config

import (
	"fmt"
	"testing"
)

func TestParseConfig(t *testing.T) {
	dsn, err := ParseConfig("demo.cnf")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dsn)
}

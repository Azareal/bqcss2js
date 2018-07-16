package main

import (
	"testing"

	"github.com/Azareal/bqcss2js/parse"
)

func failOnErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func expect(t *testing.T, success bool, msg string) {
	if !success {
		t.Fatal(msg)
	}
}

func TestParse(t *testing.T) {
	queries, err := parse.ParseBytes([]byte(""))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("hi"))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("@"))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("@element"))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")
}

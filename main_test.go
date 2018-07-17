package main

import (
	"runtime/debug"
	"testing"

	"github.com/Azareal/bqcss2js/parse"
)

func failOnErr(t *testing.T, err error) {
	if err != nil {
		debug.PrintStack()
		t.Fatal(err)
	}
}

func expect(t *testing.T, success bool, msg string) {
	if !success {
		debug.PrintStack()
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

	queries, err = parse.ParseBytes([]byte("@element "))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("@element .hi"))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("@element .hi {"))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("@element .hi {}"))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("@element .hi { }"))
	failOnErr(t, err)
	expect(t, len(queries) == 0, "The AST should be empty")

	queries, err = parse.ParseBytes([]byte("@element .hi {...}"))
	failOnErr(t, err)
	expect(t, len(queries) == 1, "The AST should have one query")

	queries, err = parse.ParseBytes([]byte(`@element .hi {
	.row {
		background-color: red;
	}
}`))
	failOnErr(t, err)
	expect(t, len(queries) == 1, "The AST should have one query")
}

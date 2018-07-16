package parse

import (
	"runtime/debug"
	"testing"
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

// peekmatch is offset up by one and if it doesn't find what it's looking for immediately there, it aborts
func TestPeekMatch(t *testing.T) {
	var match = func(cur int, phrase string, haystack string) bool {
		return peekMatch(cur, phrase, []rune(haystack))
	}
	expect(t, !match(0, "h", ""), "h shouldn't match a blank string")
	expect(t, !match(0, "hi", ""), "hi shouldn't match a blank string")
	expect(t, !match(0, "h", "h"), "h shouldn't match h")
	expect(t, !match(0, "hi", "hi"), "hi shouldn't match hi")
	expect(t, !match(0, "hi", "hih"), "hi shouldn't match hih")
	expect(t, !match(0, "hi", "hihi"), "hi shouldn't match hihi")
	expect(t, match(0, "h", " h"), "h should match  h")
	expect(t, match(0, "h", " h "), "h should match  h ")
	expect(t, match(0, "hi", " hi"), "h should match  hi")
	expect(t, match(0, "hi", " hi "), "h should match  hi ")
}

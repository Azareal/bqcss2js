package parse

import (
	"io/ioutil"
	"strings"
)

// TODO: Do some sort of streaming I/O rather than loading it all into memory?
// TODO: Write a test for this
func ParseFile(path string) error {
	path = strings.Replace(path, "\\", "/", -1)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return ParseBytes(data)
}

/*
// Sample so I don't forget what the syntax looks like
@element '.widget' and (min-width: 500px) {
	.row {
		background: red;
	}
}
// Sample so I don't forget what the syntax looks like
@element '.widget' and (min-width: 500px) {
	background: red;
}
*/
// TODO: Write a test for this
// TODO: Support single line comments? It's one of my gripes with CSS in general
func ParseBytes(data []byte) error {
	// TODO: Optimise this
	dataRunes := []rune(string(data))
	for i := 0; i < len(dataRunes); i++ {
		char := dataRunes[i]
		// Skip comments, even if they contain @
		if char == '/' && peekMatch(i, "*", dataRunes) {
			for ; i < len(dataRunes); i++ {
				char := dataRunes[i]
				if char == '*' && peekMatch(i, "/", dataRunes) {
					break
				}
			}
		} else if char == '@' && peekMatch(i, "element ", dataRunes) {
			//
		}
	}
	return nil
}

// TODO: Test this
func peekMatch(cur int, phrase string, runes []rune) bool {
	if cur+len(phrase) > len(runes) {
		return false
	}
	for i, char := range phrase {
		if runes[cur+i+1] != char {
			return false
		}
	}
	return true
}

// TODO: Test this
func tryStepForward(i int, step int, runes []rune) (int, bool) {
	i += step
	if i < len(runes) {
		return i, true
	}
	return i - step, false
}

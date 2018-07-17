package parse

import (
	"io/ioutil"
	"log"
	"strings"
)

// TODO: Do some sort of streaming I/O rather than loading it all into memory?
// TODO: Write a test for this
func ParseFile(path string) ([]*Query, error) {
	path = strings.Replace(path, "\\", "/", -1)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

/*
// Sample so I don't forget what the syntax looks like
@element .widget and (min-width: 500px) {
	.row {
		background: red;
	}
}
// TODO: Implement this short style
// Sample so I don't forget what the syntax looks like
@element .widget and (min-width: 500px) {
	background: red;
}
*/
// TODO: Write a test for this
// TODO: Support single line comments? It's one of my gripes with CSS in general
func ParseBytes(data []byte) (queries []*Query, err error) {
	// TODO: Optimise this
	dataRunes := []rune(string(data))
	for i := 0; i < len(dataRunes); i++ {
		char := dataRunes[i]
		// Skip comments, even if they contain @
		if char == '/' && peekMatch(i, "*", dataRunes) {
			i = skipCharsUntil(i, "*/", dataRunes)
		} else if char == '@' && peekMatch(i, "element ", dataRunes) {
			i += 8
			initialI := i
			query := new(Query)
			// TODO: Allow quotes around the selector like EQCSS?
			for ; i < len(dataRunes); i++ {
				char := dataRunes[i]
				// To cover :not(), :has(), etc.
				// TODO: Add multiple passes to speed up the queries which don't have ':'?
				if char == ':' && peekMatch(i, "(", dataRunes) {
					i = skipCharsUntil(i, ")", dataRunes)
					i--
				} else if char == ' ' && peekMatch(i, " and (", dataRunes) {
					query.Selector = strings.TrimSpace(string(dataRunes[initialI:i]))
					i += 6
					// TODO: We could support multiple media rules, but let's keep it simple for now
					startDelim := i
					i = skipCharsUntil(i, ")", dataRunes)
					query.MediaRules = strings.TrimSpace(string(dataRunes[startDelim:i]))
					break
				} else if char == '(' || char == '{' {
					query.Selector = strings.TrimSpace(string(dataRunes[initialI:i]))
					break
				}
			}
			if i < len(dataRunes) {
				var ok bool
				i, ok = tryStepForward(i, 1, dataRunes)
				if !ok {
					continue
				}

				initialI := i
				braceCount := 1
				for ; i < len(dataRunes); i++ {
					char := dataRunes[i]
					// TODO: Ignore braces in things like content: ""
					if char == '{' {
						braceCount++
					} else if char == '}' {
						braceCount--
					}
					if braceCount == 0 {
						break
					}
				}
				query.Body = strings.TrimSpace(string(dataRunes[initialI:i]))

				if query.Selector != "" && query.Body != "" {
					log.Printf("adding query: %+v\n", query)
					queries = append(queries, query)
				}
			}
		}
	}
	return queries, nil
}

func skipCharsUntil(i int, until string, dataRunes []rune) (newI int) {
	if len(until) == 0 {
		return i
	}
	if len(until) == 1 {
		for ; i < len(dataRunes); i++ {
			char := dataRunes[i]
			if char == rune(until[0]) {
				break
			}
		}
	} else {
		for ; i < len(dataRunes); i++ {
			char := dataRunes[i]
			if char == rune(until[0]) && peekMatch(i, until[1:], dataRunes) {
				break
			}
		}
	}
	return i
}

// TODO: Test this
func peekMatch(cur int, phrase string, runes []rune) bool {
	if cur+len(phrase) > len(runes) {
		return false
	}
	for i, char := range phrase {
		if cur+i+1 >= len(runes) {
			return false
		}
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

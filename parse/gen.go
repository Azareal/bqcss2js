package parse

import (
	"bytes"
	"strconv"
)

// TODO: Strip the newlines with an init() function to reduce the amount of data sent over the wire?
// TODO: Allow some scripts to skip attachObserver, if they opt for manual calls to widthChanged instead
var jsPrefix = []byte(`"use strict"
function widthChanged(element) {
	//
}
function attachObserver(element) {
	let observer = new MutationObserver((mutList) => {
		for(var mutation of mutList) {
			if(mutation.type=="attributes" && mutation.attributeName=="width") {
				widthChanged(element);
			}
		}
	});
	observer.observe(element, {attributes:true, childList: true, subtree: true});
}
let elemsToStyleIDs = {};
let selToStyleID = {
`)

// TODO: Add support for styling the immediate element not just the descendents
// TODO: Is ResizeObserver a good fallback for MutationObserver?
// ? - Can we use V8 / Node / Otto for testing the generated javascript? Maybe spin it up in the browser or something?
func Output(queries []*Query) (cssBytes []byte, jsBytes []byte, err error) {
	seed := make([]byte, len(jsPrefix))
	copy(seed, jsPrefix)
	jsBuf := bytes.NewBuffer(seed)
	cssBuf := bytes.NewBuffer(nil)
	for i, query := range queries {
		cssBuf.WriteString("[fairy-dust-" + strconv.Itoa(i) + "]{")
		cssBuf.WriteString(query.Body)
		cssBuf.Write([]byte("}\n"))
		jsBuf.Write([]byte(`"` + query.Selector + `"` + ":" + strconv.Itoa(i)))
	}
	jsBuf.Write([]byte("};"))
	return cssBuf.Bytes(), jsBuf.Bytes(), nil
}

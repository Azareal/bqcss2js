package parse

import (
	"bytes"
	"strconv"
)

// TODO: Add support for styling the immediate element not just the descendents
// TODO: Is ResizeObserver a good fallback for MutationObserver?
// ? - Can we use V8 / Node / Otto for testing the generated javascript? Maybe spin it up in the browser or something?
func Output(queries []*Query, config *Config) (cssBytes []byte, jsBytes []byte, err error) {
	// TODO: Strip the newlines with an init() function to reduce the amount of data sent over the wire?
	// TODO: Allow some scripts to skip attachObserver, if they opt for manual calls to widthChanged instead
	// TODO: Vary the output based on which rules are used, e.g. don't output widthChanged, if there's no width related queries
	// TODO: Use a class instead of loose JS functions
	var jsPrefix = `"use strict"
const RuleTypes = Object.freeze({"minWidth":0, "maxWidth":1});`
	// TODO: Run the autostart function on document.ready?
	if config.Autostart {
		jsPrefix += `
(() => {
	initBqcss();
})();`
	}
	jsPrefix += `
function initBqcss() {
	Object.keys(selToStyleID).forEach((key) => {
		let elements = document.querySelectorAll(key)
		for(var i = 0; i < elements.length; i++) {
			elemMappings[elements[i]] = {
				"styleID": selToStyleID[key],
			};
		}
	});
}
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
var elemMappings = {};
var selToStyleID = {
`
	jsBuf := bytes.NewBuffer([]byte(jsPrefix))
	cssBuf := bytes.NewBuffer(nil)
	for i, query := range queries {
		cssBuf.WriteString("[bqcss-" + strconv.Itoa(i) + "]{")
		cssBuf.WriteString(query.Body)
		cssBuf.Write([]byte("}\n"))
		jsBuf.WriteString(`"` + query.Selector + `"` + ":" + strconv.Itoa(i) + ",")
	}
	jsBuf.Write([]byte("};"))
	return cssBuf.Bytes(), jsBuf.Bytes(), nil
}

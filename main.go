// ! Warning: The API in this library might change frequently without warning
package main

import (
	"flag"
	"log"

	"github.com/Azareal/bqcss2js/parse"
)

// TODO: Support directories as origins
// ? - Should we allow media queries inside element queries to have element queries which are bounded to certain viewport sizes? Does this even make sense?
// ? - Might make sense to allow nested element queries? What impact would this have on performance?
func main() {
	originFilePtr := flag.String("originFile", "", "the path to the file you want to read from")
	destFilePtr := flag.String("destFile", "", "the path to the file you want to write to")
	flag.Parse()

	originFile := *originFilePtr
	destFile := *destFilePtr
	_ = destFile

	_, err := parse.ParseFile(originFile)
	if err != nil {
		log.Fatal(err)
	}
}

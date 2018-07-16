package main

import (
	"flag"
	"log"

	"github.com/Azareal/bqcss2js/parse"
)

// TODO: Support directories as origins
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

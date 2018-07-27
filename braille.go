/*
random package message
*/
package main

import (
	"flag"
	"strings"

	"./brailleFuncs"
)

func main() {
	var fn = flag.String("f", "", "Braille file to Message")
	var ms = flag.String("m", "", "Message to Braille")
	flag.Parse()

	switch {
	case *fn != "":
		brailleFuncs.BrailleToMessage(*fn)
	case *ms != "":
		brailleFuncs.MessageToBraille(strings.ToLower(*ms))
	default:
		flag.Usage()
	}
}

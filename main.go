package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	w := flag.CommandLine.Output() // stderr
	noNewline := flag.Bool("n", false, "don't print newline")
	noEncode := flag.Bool("u", false, "don't encode output")

	flag.Usage = func() {
		fmt.Fprintf(w, "Usage: %s [<args>] <json>\n", os.Args[0])
		fmt.Fprintf(w, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	b := flag.Arg(0)

	var v1 interface{}
	err := bson.UnmarshalJSON([]byte(b), &v1)
	if err != nil {
		fmt.Fprintf(w, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	v2, err := bson.Marshal(v1)
	if err != nil {
		fmt.Fprintf(w, "Error marshalling BSON: %v\n", err)
		os.Exit(1)
	}

	if *noEncode {
		fmt.Print(string(v2))
	} else {
		fmt.Print(base64.StdEncoding.EncodeToString(v2))
	}

	if !*noNewline {
		fmt.Print("\n")
	}
}

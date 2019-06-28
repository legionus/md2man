package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/russross/blackfriday.v2"
)

var (
	showHelp = flag.Bool("help", false, "show this text and exit")
	outFile  = flag.String("output", "", "send output to `FILE'")
)

func Render(doc []byte) []byte {
	renderer := RoffRenderer(0)
	return blackfriday.Run(doc,
		blackfriday.WithRenderer(renderer),
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.Titleblock),
	)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] markdown-file\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	var err error
	output := os.Stdout

	if len(*outFile) > 0 {
		output, err = os.Create(*outFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	output.Write(Render(input))
}

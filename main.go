package main

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"log"

	"github.com/russross/blackfriday"
)

var (
	showHelp = flag.Bool("help", false, "show this text and exit")
	outFile = flag.String("output", "", "send output to `FILE'")
)

func Render(doc []byte) []byte {
	renderer := RoffRenderer(0)
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_FOOTNOTES
	extensions |= blackfriday.EXTENSION_TITLEBLOCK

	return blackfriday.Markdown(doc, renderer, extensions)
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

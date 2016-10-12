package main

import (
	_ "fmt"
	"flag"
	"os"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()

	args := flag.Args()

	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(Render(input))
}

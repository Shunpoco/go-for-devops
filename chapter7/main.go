package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	useProd = flag.Bool(
		"prod",
		true,
		"Use a production endpoint",
	)
	useDev = flag.Bool(
		"dev",
		false,
		"Use a development endpoint",
	)
	help = new(bool)
)

func init() {
	flag.BoolVar(help, "help", false, "Display help text")
	flag.BoolVar(help, "h", false, "Display help text (shorthand)")
}

func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		return
	}

	authors := flag.Args()
	if len(authors) == 0 {
		log.Println("did not pass any authors")
		os.Exit(1)
	}

	for _, author := range authors {
		fmt.Println(author)
	}
}

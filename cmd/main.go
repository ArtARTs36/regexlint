package main

import (
	"flag"
	"log"

	"github.com/artarts36/regexlint/internal/linter"
	"github.com/artarts36/regexlint/internal/loader"
)

func main() {
	flag.Parse()
	syntax := requireInputString("syntax", 0)
	source := requireInputString("source", 1)
	sourcePointer := flag.Arg(2)

	l := linter.NewLinter(loader.NewChain())

	result, err := l.Lint(syntax, source, sourcePointer)
	if err != nil {
		log.Fatalln(err)
	}

	for _, regex := range result.Regexes {
		if regex.Valid() {
			log.Printf("regex %q is valid", regex.String)
		} else {
			log.Printf("regex %q is invalid", regex.String)
		}
	}

	if result.Failed() {
		log.Fatalf("found %d invalid regexes", result.Fails)
	}
}

func requireInputString(name string, order int) string {
	value := flag.Arg(order)

	if value != "" {
		return value
	}

	log.Fatalf("%q in %d argument must be set", name, order)

	return ""
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/artarts36/regexlint/internal/linter"
	"github.com/artarts36/regexlint/internal/loader"
	"github.com/artarts36/regexlint/internal/syntax"
	"github.com/artarts36/singlecli"
)

var (
	syntaxMap = syntax.CreateSyntaxMap()
)

func main() {
	app := &cli.App{
		BuildInfo: &cli.BuildInfo{
			Name:      "regexlint",
			Version:   "1.0.0",
			BuildDate: time.Now().Format(time.DateTime),
		},
		Args: []*cli.ArgDefinition{
			{
				Name:        "syntax",
				Required:    true,
				Description: "syntax of regex",
				ValuesEnum:  syntaxNames(),
			},
			{
				Name:        "source",
				Required:    true,
				Description: "source of regex (value or path to file)",
			},
			{
				Name:        "sourcePointer",
				Required:    false,
				Description: "source pointer",
			},
		},
		Action: lint,
		UsageExamples: []*cli.UsageExample{
			{
				Command:     "regexlint go file.yaml headers.cors",
				Description: "Check golang regex from headers.cors of file.yaml",
			},
			{
				Command:     "regexlint pcre file.txt row-0",
				Description: "Check PCRE from row 0 of file.txt",
			},
		},
	}

	app.RunWithGlobalArgs(context.Background())
}

func lint(ctx *cli.Context) error {
	syntaxArg := ctx.GetArg("syntax")
	source := ctx.GetArg("source")
	sourcePointer := ctx.GetArg("sourcePointer")

	l := linter.NewLinter(loader.NewChain(), syntaxMap)

	result, err := l.Lint(syntaxArg, source, sourcePointer)
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
		return fmt.Errorf("found %d invalid regexes", result.Fails)
	}

	return nil
}

func syntaxNames() []string {
	ks := make([]string, 0, len(syntaxMap))
	for k := range syntaxMap {
		ks = append(ks, k)
	}
	return ks
}

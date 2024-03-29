package main

import (
	"context"
	"fmt"
	"log"

	"github.com/artarts36/singlecli"
	"github.com/artarts36/singlecli/color"

	"github.com/artarts36/regexlint/internal/linter"
	"github.com/artarts36/regexlint/internal/loader"
	"github.com/artarts36/regexlint/internal/syntax"
)

var (
	Version   = "dev"
	BuildDate = "2024-02-22 22:22:22"

	syntaxMap = syntax.CreateSyntaxMap()
)

func main() {
	app := &cli.App{
		BuildInfo: &cli.BuildInfo{
			Name:      "regexlint",
			Version:   Version,
			BuildDate: BuildDate,
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
				Command:     "regexlint go file.json headers.cors",
				Description: "Check golang regex from headers.cors of file.json",
			},
			{
				Command:     "regexlint pcre file.txt row-0",
				Description: "Check PCRE from row 0 of file.txt",
			},
			{
				Command:     "regexlint pcre file.txt row-all",
				Description: "Check PCRE from each row of file.txt",
			},
			{
				Command:     "regexlint pcre \"php|golang\"",
				Description: "Check PCRE from command argument",
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
			log.Println(color.Green("regex %q is valid", regex.String))
		} else {
			log.Println(color.Red("regex %q is invalid", regex.String))
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

package linter

import (
	"fmt"

	"github.com/artarts36/regexlint/internal"
	"github.com/artarts36/regexlint/internal/syntax"
)

type Linter struct {
	syntax       map[string]Syntax
	sourceLoader regexLoader
}

type Syntax interface {
	Lint(regex string) (*internal.Regex, error)
}

type regexLoader interface {
	Supports(source, pointer string) bool
	Load(source, pointer string) ([]string, error)
}

func NewLinter(regexLoader regexLoader) *Linter {
	goSyntax := &syntax.Go{}
	pcreSyntax := &syntax.PCRE{}

	return &Linter{
		syntax: map[string]Syntax{
			"go":     goSyntax,
			"golang": goSyntax,
			"pcre":   pcreSyntax,
			"php":    pcreSyntax,
			"perl":   pcreSyntax,
		},
		sourceLoader: regexLoader,
	}
}

func (l *Linter) Lint(lang, source, sourcePointer string) (*LintResult, error) {
	regexes, err := l.sourceLoader.Load(source, sourcePointer)
	if err != nil {
		return nil, fmt.Errorf("unable to load regexes: %s", err)
	}

	s, found := l.syntax[lang]
	if !found {
		return nil, fmt.Errorf("syntax %q not found", lang)
	}

	result := &LintResult{
		Regexes: make([]*internal.Regex, 0, len(regexes)),
	}

	for _, regex := range regexes {
		iRegex, lintErr := s.Lint(regex)
		if lintErr != nil {
			return nil, fmt.Errorf("lint failed: %s", err)
		}

		if !iRegex.Valid() {
			result.Fails++
		}

		result.Regexes = append(result.Regexes, iRegex)
	}

	return result, nil
}

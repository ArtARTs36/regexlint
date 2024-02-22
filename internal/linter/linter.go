package linter

import (
	"fmt"

	"github.com/artarts36/regexlint/internal"
	"github.com/artarts36/regexlint/internal/syntax"
)

type Linter struct {
	syntax       map[string]syntax.Syntax
	sourceLoader regexLoader
}

type regexLoader interface {
	Load(source *internal.RegexSource, pointer string) ([]string, error)
}

func NewLinter(regexLoader regexLoader, syntax map[string]syntax.Syntax) *Linter {
	return &Linter{
		syntax:       syntax,
		sourceLoader: regexLoader,
	}
}

func (l *Linter) Lint(lang, source, sourcePointer string) (*LintResult, error) {
	rSource := internal.NewRegexSource(source)

	regexes, err := l.sourceLoader.Load(rSource, sourcePointer)
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

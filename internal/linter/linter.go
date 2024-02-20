package linter

import (
	"fmt"

	"github.com/artarts36/regexlint/internal"
	"github.com/artarts36/regexlint/internal/syntax"
)

type Linter struct {
	syntax       map[string]Syntax
	sourceLoader SourceLoader
}

type Syntax interface {
	Lint(regex string) (*internal.Regex, error)
}

type SourceLoader interface {
	Supports(source, pointer string) bool
	Load(source, pointer string) (string, error)
}

func NewLinter(sourceLoader SourceLoader) *Linter {
	goSyntax := &syntax.Go{}

	return &Linter{
		syntax: map[string]Syntax{
			syntax.GoName:      goSyntax,
			syntax.GoAliasName: goSyntax,
		},
		sourceLoader: sourceLoader,
	}
}

func (l *Linter) Lint(lang, source, sourcePointer string) (*internal.Regex, error) {
	regex, err := l.sourceLoader.Load(source, sourcePointer)
	if err != nil {
		return nil, fmt.Errorf("unable to load regex: %s", err)
	}

	s, found := l.syntax[lang]
	if !found {
		return nil, fmt.Errorf("syntax %q not found", lang)
	}

	iRegex, err := s.Lint(regex)
	if err != nil {
		return iRegex, fmt.Errorf("source invalid: %s", err.Error())
	}

	return iRegex, nil
}

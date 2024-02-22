package syntax

import "github.com/artarts36/regexlint/internal"

type Syntax interface {
	Lint(regex string) (*internal.Regex, error)
}

func CreateSyntaxMap() map[string]Syntax {
	goSyntax := &Go{}
	pcreSyntax := &PCRE{}

	return map[string]Syntax{
		"go":     goSyntax,
		"golang": goSyntax,
		"pcre":   pcreSyntax,
		"php":    pcreSyntax,
		"perl":   pcreSyntax,
	}
}

package linter

import "github.com/artarts36/regexlint/internal"

type LintResult struct {
	Regexes []*internal.Regex
	Fails   int
}

func (r *LintResult) Failed() bool {
	return r.Fails > 0
}

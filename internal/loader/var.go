package loader

import "github.com/artarts36/regexlint/internal"

type Var struct {
}

func (v *Var) Supports(source *internal.RegexSource, sourcePointer string) bool {
	return !source.IsFile() && sourcePointer == ""
}

func (v *Var) Load(source *internal.RegexSource, _ string) ([]string, error) {
	return []string{source.Source}, nil
}

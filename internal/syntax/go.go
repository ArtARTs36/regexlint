package syntax

import (
	"github.com/artarts36/regexlint/internal"
	"regexp"
)

const (
	GoName      = "go"
	GoAliasName = "golang"
)

type Go struct {
}

func (s *Go) Lint(regex string) (*internal.Regex, error) {
	_, err := regexp.Compile(regex)
	if err != nil {
		return &internal.Regex{
			String: regex,
		}, err
	}

	return &internal.Regex{
		String: regex,
	}, err
}

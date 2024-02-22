package syntax

import (
	"regexp"

	"github.com/artarts36/regexlint/internal"
)

type Go struct {
}

func (s *Go) Lint(regex string) (*internal.Regex, error) {
	_, err := regexp.Compile(regex)
	if err != nil {
		return &internal.Regex{
			String: regex,
			Error:  err,
		}, nil
	}

	return &internal.Regex{
		String: regex,
	}, nil
}

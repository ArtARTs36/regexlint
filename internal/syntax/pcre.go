package syntax

import (
	pcre "github.com/GRbit/go-pcre"
	"github.com/artarts36/regexlint/internal"
)

type PCRE struct {
}

func (s *PCRE) Lint(regex string) (*internal.Regex, error) {
	_, err := pcre.CompileParse(regex) //nolint:nilerr // it's contract
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

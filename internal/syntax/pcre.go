package syntax

import (
	pcre "github.com/GRbit/go-pcre" //nolint:typecheck // not advisable
	"github.com/artarts36/regexlint/internal"
)

type PCRE struct {
}

func (s *PCRE) Lint(regex string) (*internal.Regex, error) {
	_, err := pcre.CompileParse(regex)
	if err != nil {
		return &internal.Regex{
			String: regex,
		}, err
	}

	return &internal.Regex{
		String: regex,
	}, err
}

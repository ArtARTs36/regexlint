package loader_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/artarts36/regexlint/internal/loader"
)

func TestTxtAll_Load(t *testing.T) {
	l := &loader.TxtAll{}

	cases := []struct {
		Name            string
		Path            string
		ExpectedRegexes []string
		ExpectedErr     error
	}{
		{
			Name:        "fail on empty file",
			Path:        "./testdata/txt-all/empty.txt",
			ExpectedErr: errors.New("regexes not found in: ./testdata/txt-all/empty.txt"),
		},
		{
			Name: "load 2 regex",
			Path: "./testdata/txt-all/two_regexes.txt",
			ExpectedRegexes: []string{
				"string1",
				"string2",
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Name, func(t *testing.T) {
			regexes, err := l.Load(tCase.Path, "")
			if tCase.ExpectedErr != nil {
				assert.Equal(t, tCase.ExpectedErr, err)
			} else {
				require.NoError(t, tCase.ExpectedErr)

				assert.Equal(t, tCase.ExpectedRegexes, regexes)
			}
		})
	}
}

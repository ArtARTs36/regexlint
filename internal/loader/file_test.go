package loader_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/artarts36/regexlint/internal"
	"github.com/artarts36/regexlint/internal/loader"
)

func TestFile_Supports(t *testing.T) {
	l := &loader.UnmarshallingFile{}

	cases := []struct {
		Name       string
		Path       string
		Extensions []string
		Expected   bool
	}{
		{
			Name:       "false on empty extensions",
			Path:       "./testdata/txt-all/empty.txt",
			Extensions: []string{},
		},
		{
			Name: "false on source not contains need extension",
			Path: "./testdata/txt-all/two_regexes.txt",
			Extensions: []string{
				"yaml",
			},
		},
		{
			Name: "ok",
			Path: "./testdata/txt-all/two_regexes.txt",
			Extensions: []string{
				"txt",
			},
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.Name, func(t *testing.T) {
			got := l.Supports(internal.NewRegexSource(tCase.Path), "")

			assert.Equal(t, tCase.Expected, got)
		})
	}
}

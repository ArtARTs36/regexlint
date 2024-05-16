package internal_test

import (
	"fmt"
	"testing"

	"github.com/artarts36/regexlint/internal"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRegexSource(t *testing.T) {
	cases := []struct {
		Source         string
		ExpectedSource string
		ExpectedExt    string
	}{
		{
			Source:         "invalid-regex",
			ExpectedSource: "invalid-regex",
		},
		{
			Source:         "invalid-regex/.",
			ExpectedSource: "invalid-regex/.",
		},
		{
			Source:         "invalid-regex/.yaml",
			ExpectedSource: "invalid-regex/.yaml",
		},
		{
			Source:         "regex_test.go",
			ExpectedSource: "regex_test.go",
			ExpectedExt:    "go",
		},
	}

	for i, tCase := range cases {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			r := internal.NewRegexSource(tCase.Source)

			require.Equal(t, tCase.ExpectedSource, r.Source)
			assert.True(t, r.HasFileExtension(tCase.ExpectedExt))
		})
	}
}

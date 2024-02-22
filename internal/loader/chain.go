package loader

import (
	"fmt"

	"github.com/artarts36/regexlint/internal"
)

type Chain struct {
	loaders []loader
}

type loader interface {
	Supports(source *internal.RegexSource, sourcePointer string) bool
	Load(source *internal.RegexSource, _ string) ([]string, error)
}

func (c *Chain) Supports(source *internal.RegexSource, sourcePointer string) bool {
	for _, l := range c.loaders {
		if l.Supports(source, sourcePointer) {
			return true
		}
	}

	return false
}

func (c *Chain) Load(source *internal.RegexSource, sourcePointer string) ([]string, error) {
	for _, l := range c.loaders {
		if l.Supports(source, sourcePointer) {
			return l.Load(source, sourcePointer)
		}
	}

	return []string{}, fmt.Errorf("source loader not found")
}

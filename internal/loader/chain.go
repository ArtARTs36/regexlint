package loader

import "fmt"

type Chain struct {
	loaders []loader
}

type loader interface {
	Supports(source, sourcePointer string) bool
	Load(source, _ string) (string, error)
}

func (c *Chain) Supports(source, sourcePointer string) bool {
	for _, l := range c.loaders {
		if l.Supports(source, sourcePointer) {
			return true
		}
	}

	return false
}

func (c *Chain) Load(source, sourcePointer string) (string, error) {
	for _, l := range c.loaders {
		if l.Supports(source, sourcePointer) {
			return l.Load(source, sourcePointer)
		}
	}

	return "", fmt.Errorf("source loader not found")
}

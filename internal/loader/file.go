package loader

import (
	"fmt"
	"os"

	"github.com/artarts36/regexlint/internal"
	"github.com/artarts36/regexlint/internal/dot"
)

type UnmarshallingFile struct {
	extensions  []string
	unmarshaler fileUnmarshaler
}

type fileUnmarshaler func(content []byte, out interface{}) error

func (f *UnmarshallingFile) Supports(source *internal.RegexSource, sourcePointer string) bool {
	return sourcePointer != "" && f.supportsExtensions(source)
}

func (f *UnmarshallingFile) Load(source *internal.RegexSource, pointer string) ([]string, error) {
	file, err := os.ReadFile(source.Source)
	if err != nil {
		return []string{}, fmt.Errorf("unable to read file: %s", err)
	}

	var val map[string]interface{}

	err = f.unmarshaler(file, &val)
	if err != nil {
		return []string{}, fmt.Errorf("unable to unmarshal yaml: %s", err)
	}

	pointers := splitPointer(pointer)
	regexes := make([]string, 0, len(pointers))
	for _, p := range pointers {
		s, dotErr := dot.FindString(val, p)
		if dotErr != nil {
			return []string{}, dotErr
		}

		regexes = append(regexes, s)
	}

	return regexes, nil
}

func (f *UnmarshallingFile) supportsExtensions(source *internal.RegexSource) bool {
	for _, extension := range f.extensions {
		if source.HasFileExtension(extension) {
			return true
		}
	}

	return false
}

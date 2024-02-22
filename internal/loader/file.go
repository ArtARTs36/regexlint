package loader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/artarts36/regexlint/internal/dot"
)

type UnmarshallingFile struct {
	extensions  []string
	unmarshaler fileUnmarshaler
}

type fileUnmarshaler func(content []byte, out interface{}) error

func (f *UnmarshallingFile) Supports(source, sourcePointer string) bool {
	return sourcePointer != "" && f.supportsExtensions(source)
}

func (f *UnmarshallingFile) Load(source, pointer string) ([]string, error) {
	file, err := os.ReadFile(source)
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

func (f *UnmarshallingFile) supportsExtensions(source string) bool {
	ext := filepath.Ext(source)
	if ext == "" {
		return false
	}

	for _, extension := range f.extensions {
		if extension == ext {
			return true
		}
	}

	return false
}

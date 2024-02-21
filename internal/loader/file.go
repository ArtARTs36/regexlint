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

func (y *UnmarshallingFile) Supports(source, sourcePointer string) bool {
	return sourcePointer != "" && y.supportsExtensions(source)
}

func (y *UnmarshallingFile) Load(source, pointer string) ([]string, error) {
	file, err := os.ReadFile(source)
	if err != nil {
		return []string{}, fmt.Errorf("unable to read file: %s", err)
	}

	var val map[string]interface{}

	err = y.unmarshaler(file, &val)
	if err != nil {
		return []string{}, fmt.Errorf("unable to unmarshal yaml: %s", err)
	}

	s, err := dot.FindString(val, pointer)
	if err != nil {
		return []string{}, err
	}

	return []string{s}, nil
}

func (y *UnmarshallingFile) supportsExtensions(source string) bool {
	ext := filepath.Ext(source)
	if ext == "" {
		return false
	}

	for _, extension := range y.extensions {
		if extension == ext {
			return true
		}
	}

	return true
}

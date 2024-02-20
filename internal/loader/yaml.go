package loader

import (
	"fmt"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v3"

	"github.com/artarts36/regexlint/internal/dot"
)

type YAML struct {
}

func (y *YAML) Supports(source, sourcePointer string) bool {
	return sourcePointer != "" && (strings.HasSuffix(source, ".yaml") || strings.HasSuffix(source, ".yml"))
}

func (y *YAML) Load(source, pointer string) (string, error) {
	file, err := os.ReadFile(source)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %s", err)
	}

	var val map[string]interface{}

	err = yaml.Unmarshal(file, &val)
	if err != nil {
		return "", fmt.Errorf("unable to unmarshal yaml: %s", err)
	}

	s, err := dot.FindString(val, pointer)
	if err != nil {
		return "", err
	}

	return s, nil
}

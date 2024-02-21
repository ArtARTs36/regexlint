package loader

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/artarts36/regexlint/internal/dot"
)

type JSON struct {
}

func (y *JSON) Supports(source, sourcePointer string) bool {
	return sourcePointer != "" && strings.HasSuffix(source, ".json")
}

func (y *JSON) Load(source, pointer string) ([]string, error) {
	file, err := os.ReadFile(source)
	if err != nil {
		return []string{}, fmt.Errorf("unable to read file: %s", err)
	}

	var val map[string]interface{}

	err = json.Unmarshal(file, &val)
	if err != nil {
		return []string{}, fmt.Errorf("unable to unmarshal yaml: %s", err)
	}

	s, err := dot.FindString(val, pointer)
	if err != nil {
		return []string{}, err
	}

	return []string{s}, nil
}

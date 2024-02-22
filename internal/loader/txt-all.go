package loader

import (
	"fmt"
	"os"

	"strings"
)

type TxtAll struct {
}

func (y *TxtAll) Supports(source, sourcePointer string) bool {
	return strings.HasSuffix(source, ".txt") && sourcePointer == "row-all"
}

func (y *TxtAll) Load(source, _ string) ([]string, error) {
	file, err := os.ReadFile(source)
	if err != nil {
		return []string{}, fmt.Errorf("unable to read file: %s", err)
	}

	regexes := strings.SplitAfter(string(file), "\n")
	if len(regexes) == 0 {
		return []string{}, fmt.Errorf("regexes not found in: %s", file)
	}

	if regexes[len(regexes)-1] == "" {
		regexes = regexes[:len(regexes)-1]
	}

	return regexes, nil
}

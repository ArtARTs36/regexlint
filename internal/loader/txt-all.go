package loader

import (
	"fmt"
	"os"

	"strings"
)

type TxtAll struct {
}

func (y *TxtAll) Supports(source, sourcePointer string) bool {
	return strings.HasSuffix(source, ".txt-all") && sourcePointer == "row-all"
}

func (y *TxtAll) Load(source, _ string) ([]string, error) {
	file, err := os.ReadFile(source)
	if err != nil {
		return []string{}, fmt.Errorf("unable to read file: %s", err)
	}

	regexes := strings.Split(string(file), "\n")
	if len(regexes) == 0 || (len(regexes) == 1 && y.lastRowIsEmpty(regexes)) {
		return []string{}, fmt.Errorf("regexes not found in: %s", source)
	}

	if y.lastRowIsEmpty(regexes) {
		regexes = regexes[:len(regexes)-1]
	}

	return regexes, nil
}

func (y *TxtAll) lastRowIsEmpty(regexes []string) bool {
	return regexes[len(regexes)-1] == ""
}

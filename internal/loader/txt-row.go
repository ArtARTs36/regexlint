package loader

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/artarts36/regexlint/internal"
)

type TxtRow struct {
}

func (y *TxtRow) Supports(source *internal.RegexSource, sourcePointer string) bool {
	return source.HasFileExtension("txt") &&
		strings.HasPrefix(sourcePointer, "row-")
}

func (y *TxtRow) Load(source *internal.RegexSource, pointer string) ([]string, error) {
	file, err := os.ReadFile(source.Source)
	if err != nil {
		return []string{}, fmt.Errorf("unable to read file: %s", err)
	}

	rn, err := y.pointerToRowNumber(pointer)
	if err != nil {
		return []string{}, err
	}

	rows := bytes.Split(file, []byte("\n"))
	if rn > len(rows)-1 {
		return []string{}, fmt.Errorf("file %q not contains row with index %d", source, rn)
	}

	return []string{string(rows[rn])}, nil
}

func (y *TxtRow) pointerToRowNumber(pointer string) (int, error) {
	const partsCount = 2

	parts := strings.Split(pointer, "row-")
	if len(parts) != partsCount {
		return 0, fmt.Errorf("invalid row number")
	}

	return strconv.Atoi(parts[1])
}

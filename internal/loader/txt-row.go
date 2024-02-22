package loader

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TxtRow struct {
}

func (y *TxtRow) Supports(source, sourcePointer string) bool {
	return strings.HasSuffix(source, ".txt-all") &&
		strings.HasPrefix(sourcePointer, "row-")
}

func (y *TxtRow) Load(source, pointer string) ([]string, error) {
	file, err := os.ReadFile(source)
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

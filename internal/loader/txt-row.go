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
	return strings.HasSuffix(source, ".txt") &&
		strings.HasPrefix(sourcePointer, "row-")
}

func (y *TxtRow) Load(source, pointer string) (string, error) {
	file, err := os.ReadFile(source)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %s", err)
	}

	rn, err := y.pointerToRowNumber(pointer)
	if err != nil {
		return "", err
	}

	rows := bytes.Split(file, []byte("\n"))
	if rn > len(rows)-1 {
		return "", fmt.Errorf("file %q not contains row with index %d", source, rn)
	}

	return string(rows[rn]), nil
}

func (y *TxtRow) pointerToRowNumber(pointer string) (int, error) {
	parts := strings.Split(pointer, "row-")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid row number")
	}

	return strconv.Atoi(parts[1])
}

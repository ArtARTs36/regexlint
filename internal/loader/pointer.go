package loader

import "strings"

func splitPointer(pointer string) []string {
	return strings.Split(pointer, ",")
}

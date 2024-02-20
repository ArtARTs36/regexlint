package dot

import (
	"fmt"
	"reflect"
	"strings"
)

func FindString(val map[string]interface{}, pointer string) (string, error) {
	curr := val
	subs := strings.Split(pointer, ".")

	currentKey := ""

	for i, sub := range subs {
		if currentKey != "" {
			currentKey += "." + sub
		} else {
			currentKey += sub
		}

		item, exists := curr[sub]
		if !exists {
			return "", fmt.Errorf("key %q not found", currentKey)
		}

		v := reflect.ValueOf(item)
		switch v.Kind() { //nolint:exhaustive // no need to check all cases
		case reflect.String:
			if i == len(subs)-1 {
				return v.String(), nil
			}

			return "", fmt.Errorf("key %q must be map", currentKey)
		case reflect.Map:
			if i == len(subs)-1 {
				return "", fmt.Errorf("key %q must be string", currentKey)
			}

			c, isMap := item.(map[string]interface{})
			if !isMap {
				return "", fmt.Errorf("key %q must be map[string][]", currentKey)
			}

			curr = c
		default:
			return "", fmt.Errorf("key %q has unexpected type %q", currentKey, v.Type().Name())
		}
	}

	return "", fmt.Errorf("key %q not found", pointer)
}

package dot

import (
	"fmt"
	"reflect"
	"strings"
)

func FindString(val map[string]interface{}, pointer string) (string, error) {
	curr := val
	subs := strings.Split(pointer, ".")

	fullKey := ""

	for i, sub := range subs {
		if fullKey != "" {
			fullKey += "." + sub
		} else {
			fullKey += sub
		}

		item, exists := curr[sub]
		if !exists {
			return "", fmt.Errorf("key %q not found", fullKey)
		}

		v := reflect.ValueOf(item)
		switch v.Kind() {
		case reflect.String:
			if i == len(subs)-1 {
				return v.String(), nil
			}

			return "", fmt.Errorf("key %q must be map", fullKey)
		case reflect.Map:
			if i == len(subs)-1 {
				return "", fmt.Errorf("key %q must be string", fullKey)
			}

			curr = item.(map[string]interface{})
		default:
			return "", fmt.Errorf("key %q has unexpected type %q", fullKey, v.Type().Name())
		}
	}

	return "", fmt.Errorf("key %q not found", pointer)
}

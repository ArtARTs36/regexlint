package loader

import (
	"encoding/json"
)

func NewJSON() *UnmarshallingFile {
	return &UnmarshallingFile{
		unmarshaler: json.Unmarshal,
	}
}

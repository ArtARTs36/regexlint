package loader

import "gopkg.in/yaml.v3"

func NewYAML() *UnmarshallingFile {
	return &UnmarshallingFile{
		unmarshaler: yaml.Unmarshal,
	}
}

package internal

import (
	"os"

	"path/filepath"
)

type Regex struct {
	String string
	Error  error
}

func (r *Regex) Valid() bool {
	return r.Error == nil
}

type RegexSource struct {
	Source        string
	fileExtension string
}

func NewRegexSource(source string) *RegexSource {
	ext := filepath.Ext(source)
	if len(ext) > 0 && ext[0] == '.' {
		if _, err := os.Stat(source); err == nil {
			ext = ext[1:]
		} else {
			ext = ""
		}
	}

	return &RegexSource{
		Source:        source,
		fileExtension: ext,
	}
}

func (s *RegexSource) IsFile() bool {
	return s.fileExtension != ""
}

func (s *RegexSource) HasFileExtension(ext string) bool {
	return s.fileExtension == ext
}

package loader

type Var struct {
}

func (v *Var) Supports(source, sourcePointer string) bool {
	return source != "" && sourcePointer == ""
}

func (v *Var) Load(source, _ string) ([]string, error) {
	return []string{source}, nil
}

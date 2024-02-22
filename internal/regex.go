package internal

type Regex struct {
	String string
	Error  error
}

func (r *Regex) Valid() bool {
	return r.Error == nil
}

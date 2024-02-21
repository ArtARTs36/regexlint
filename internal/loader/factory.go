package loader

func NewChain() *Chain {
	return &Chain{
		loaders: []loader{
			NewYAML(),
			NewJSON(),
			&TxtRow{},
			&Var{},
		},
	}
}

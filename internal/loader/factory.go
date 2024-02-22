package loader

func NewChain() *Chain {
	return &Chain{
		loaders: []loader{
			NewYAML(),
			NewJSON(),
			&TxtAll{},
			&TxtRow{},
			&Var{},
		},
	}
}

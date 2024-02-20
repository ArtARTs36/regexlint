package loader

func NewChain() *Chain {
	return &Chain{
		loaders: []loader{
			&YAML{},
			&JSON{},
			&TxtRow{},
			&Var{},
		},
	}
}

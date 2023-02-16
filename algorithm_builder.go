package songci

type algorithmBuilder struct {
	*builder

	name      string
	algorithm *algorithm
}

func newAlgorithmBuilder(builder *builder, algorithm *algorithm) *algorithmBuilder {
	return &algorithmBuilder{
		builder: builder,

		algorithm: algorithm,
	}
}

func (ab *algorithmBuilder) Name(name string) *algorithmBuilder {
	ab.algorithm.scheme = name

	return ab
}

package songci

type algorithmBuilder struct {
	*builder

	algorithm *algorithm
}

func newAlgorithmBuilder(builder *builder, algorithm *algorithm) *algorithmBuilder {
	return &algorithmBuilder{
		builder: builder,

		algorithm: algorithm,
	}
}

func (ab *algorithmBuilder) Scheme(scheme string) *algorithmBuilder {
	ab.algorithm.scheme = scheme

	return ab
}

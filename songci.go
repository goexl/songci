package songci

// Songci 对外主接口
type Songci struct {
	params *params
}

func newSongci(params *params) *Songci {
	return &Songci{
		params: params,
	}
}

func (s *Songci) Verifier(getter getter) *verifierBuilder {
	return newVerifierBuilder(s.params, getter)
}

func (s *Songci) Singer(id string, getter getter) (sb *signerBuilder) {
	s.params.id = id
	sb = newSignerBuilder(s.params, getter)

	return
}

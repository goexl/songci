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

func (s *Songci) Verifier(credential string) *verifierBuilder {
	return newVerifierBuilder(s.params, credential)
}

func (s *Songci) Maker(credential string) *makerBuilder {
	return newMakerBuilder(s.params, credential)
}

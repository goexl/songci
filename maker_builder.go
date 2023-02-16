package songci

type makerBuilder struct {
	params     *params
	credential string
}

func newMakerBuilder(params *params, credential string) *makerBuilder {
	return &makerBuilder{
		params:     params,
		credential: credential,
	}
}

func (mb *makerBuilder) Zinan() *zinanBuilder {
	return newZinanBuilder(mb.params, mb.credential)
}

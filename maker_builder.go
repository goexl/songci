package songci

type makerBuilder struct {
	params *params
	id     string
	getter getter
}

func newMakerBuilder(params *params, id string, getter getter) *makerBuilder {
	return &makerBuilder{
		params: params,
		id:     id,
		getter: getter,
	}
}

func (mb *makerBuilder) Zinan() *zinanBuilder {
	return newZinanBuilder(mb.params, mb.getter)
}

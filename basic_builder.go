package songci

type basicBuilder struct {
	params *params
	core   *coreParams
	self   *basicParams
	getter getter
}

func newBasicBuilder(params *params, core *coreParams, getter getter) *basicBuilder {
	return &basicBuilder{
		params: params,
		core:   core,
		self:   newBasicParams(),
		getter: getter,
	}
}

func (bb *basicBuilder) Password(password string) *basicBuilder {
	bb.self.password = password

	return bb
}

func (bb *basicBuilder) Build() *basic {
	return newBasic(bb.params, bb.core, bb.self, bb.getter)
}

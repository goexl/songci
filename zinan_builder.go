package songci

type zinanBuilder struct {
	params *params
	core   *coreParams
	self   *zinanParams
	getter getter
}

func newZinanBuilder(params *params, core *coreParams, getter getter) *zinanBuilder {
	return &zinanBuilder{
		params: params,
		core:   core,
		self:   newZinanParams(params, core),
		getter: getter,
	}
}

func (zb *zinanBuilder) Signed(headers ...string) *zinanBuilder {
	zb.self.signed = append(zb.self.signed, headers...)

	return zb
}

func (zb *zinanBuilder) Build() *signer {
	return newSigner(zb.core, newZinan(zb.params, zb.core, zb.self, zb.getter))
}

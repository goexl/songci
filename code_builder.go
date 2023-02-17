package songci

type codeBuilder struct {
	core    *coreParams
	params  *codeParams
	authFun authFun
}

func newCodeBuilder(core *coreParams, authFun authFun) *codeBuilder {
	return &codeBuilder{
		core:    core,
		params:  newCodeParams(),
		authFun: authFun,
	}
}

func (cb *codeBuilder) Http() *codeBuilder {
	cb.params.scheme = http

	return cb
}

func (cb *codeBuilder) Build() *coder {
	return newCoder(cb.core, cb.params, cb.authFun)
}

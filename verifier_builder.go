package songci

import (
	"strings"
)

type verifierBuilder struct {
	params *params
	core   *coreParams
	getter getter
}

func newVerifierBuilder(params *params, getter getter) *verifierBuilder {
	return &verifierBuilder{
		params: params,
		core:   newCoreParams(),
		getter: getter,
	}
}

func (vb *verifierBuilder) Method(method string) *verifierBuilder {
	vb.core.method = strings.ToUpper(method)

	return vb
}

func (vb *verifierBuilder) Get() *verifierBuilder {
	vb.core.method = methodGet

	return vb
}

func (vb *verifierBuilder) Post() *verifierBuilder {
	vb.core.method = methodPost

	return vb
}

func (vb *verifierBuilder) Uri(uri string) *verifierBuilder {
	vb.core.uri = uri

	return vb
}

func (vb *verifierBuilder) Header(key string, value string) *verifierBuilder {
	vb.core.headers[key] = value

	return vb
}

func (vb *verifierBuilder) Headers(headers map[string]string) *verifierBuilder {
	vb.core.headers = headers

	return vb
}

func (vb *verifierBuilder) Payload(payload []byte) *verifierBuilder {
	vb.core.payload = payload

	return vb
}

func (vb *verifierBuilder) Build() *verifier {
	return newVerifier(vb.params, vb.core, vb.getter)
}

package songci

import (
	"strings"
)

type verifierBuilder struct {
	params *params
	self   *verifierParams
}

func newVerifierBuilder(params *params, credential string) *verifierBuilder {
	return &verifierBuilder{
		params: params,
		self:   newVerifierParams(credential),
	}
}

func (vb *verifierBuilder) Method(method string) *verifierBuilder {
	vb.self.method = strings.ToUpper(method)

	return vb
}

func (vb *verifierBuilder) Get() *verifierBuilder {
	vb.self.method = methodGet

	return vb
}

func (vb *verifierBuilder) Post() *verifierBuilder {
	vb.self.method = methodPost

	return vb
}

func (vb *verifierBuilder) Uri(uri string) *verifierBuilder {
	vb.self.method = methodGet
	values := strings.Split(uri, interrogation)
	vb.self.uri = values[0]
	if 2 == len(values) {
		vb.self.query = values[1]
	}

	return vb
}

func (vb *verifierBuilder) Header(key string, value string) *verifierBuilder {
	vb.self.headers[key] = value

	return vb
}

func (vb *verifierBuilder) Headers(headers map[string]string) *verifierBuilder {
	vb.self.headers = headers

	return vb
}

func (vb *verifierBuilder) Payload(payload []byte) *verifierBuilder {
	vb.self.payload = payload

	return vb
}

func (vb *verifierBuilder) Build() *verifier {
	return newVerifier(vb.params, vb.self)
}

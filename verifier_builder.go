package songci

import (
	"strings"
)

type verifierBuilder struct {
	token  string
	params *params
	self   *verifierParams
}

func newVerifierBuilder(token string, params *params, credential string) *verifierBuilder {
	return &verifierBuilder{
		token:  token,
		params: params,
		self:   newVerifierParams(credential),
	}
}

func (vb *verifierBuilder) method(method string) *verifierBuilder {
	vb.self.method = strings.ToUpper(method)

	return vb
}

func (vb *verifierBuilder) Uri(uri string) *verifierBuilder {
	vb.self.method = methodGet
	splits := strings.Split(uri, interrogation)
	vb.self.uri = splits[0]
	vb.self.query = splits[1]

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
	return newVerifier(vb.token, vb.params, vb.self)
}

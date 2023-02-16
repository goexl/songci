package songci

import (
	"strings"
)

type zinanBuilder struct {
	params *params
	self   *zinanParams
}

func newZinanBuilder(params *params, credential string) *zinanBuilder {
	return &zinanBuilder{
		params: params,
		self:   newZinanParams(newVerifierParams(credential)),
	}
}

func (zb *zinanBuilder) Get(uri string) *zinanBuilder {
	zb.self.method = methodGet
	values := strings.Split(uri, interrogation)
	zb.self.uri = values[0]
	if 2 == len(values) {
		zb.self.query = values[1]
	}

	return zb
}

func (zb *zinanBuilder) Post(uri string) *zinanBuilder {
	zb.self.method = methodPost
	zb.self.uri = uri

	return zb
}

func (zb *zinanBuilder) Header(key string, value string) *zinanBuilder {
	zb.self.verifierParams.headers[key] = value

	return zb
}

func (zb *zinanBuilder) Headers(headers map[string]string) *zinanBuilder {
	zb.self.verifierParams.headers = headers

	return zb
}

func (zb *zinanBuilder) Signed(headers ...string) *zinanBuilder {
	zb.self._signed = append(zb.self._signed, headers...)

	return zb
}

func (zb *zinanBuilder) Payload(payload []byte) *zinanBuilder {
	zb.self.payload = payload

	return zb
}

func (zb *zinanBuilder) Build() *maker {
	return newMaker(newZinan(zb.params, zb.self))
}

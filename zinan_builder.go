package songci

import (
	"strings"
)

type zinanBuilder struct {
	params *zinanParams
}

func newZinanBuilder() *zinanBuilder {
	return &zinanBuilder{
		params: newZinanParams(),
	}
}

func (zb *zinanBuilder) Get(uri string) *zinanBuilder {
	zb.params.method = methodGet
	splits := strings.Split(uri, interrogation)
	zb.params.uri = splits[0]
	zb.params.query = splits[1]

	return zb
}

func (zb *zinanBuilder) Post(uri string) *zinanBuilder {
	zb.params.method = methodPost
	zb.params.uri = uri

	return zb
}

func (zb *zinanBuilder) Header(key string, value string) *zinanBuilder {
	zb.params.headers[key] = value

	return zb
}

func (zb *zinanBuilder) Headers(headers map[string]string) *zinanBuilder {
	zb.params.headers = headers

	return zb
}

func (zb *zinanBuilder) Payload(payload []byte) *zinanBuilder {
	zb.params.payload = payload

	return zb
}

func (zb *zinanBuilder) Verifier() Verifier {
	return newZinanVerifier(zb.params)
}

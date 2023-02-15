package songci

type zinanBuilder struct {
	params *zinanParams
}

func newZinanBuilder() *zinanBuilder {
	return &zinanBuilder{
		params: newZinanParams(),
	}
}

func (zb *zinanBuilder) Get() *zinanBuilder {
	zb.params.method = methodGet

	return zb
}

func (zb *zinanBuilder) Post() *zinanBuilder {
	zb.params.method = methodPost

	return zb
}

func (zb *zinanBuilder) Uri(uri string) *zinanBuilder {
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

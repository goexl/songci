package songci

type signerBuilder struct {
	params *params
	core   *coreParams
	getter getter
}

func newSignerBuilder(params *params, getter getter) *signerBuilder {
	return &signerBuilder{
		params: params,
		core:   newCoreParams(),
		getter: getter,
	}
}

func (sb *signerBuilder) Host(host string) *signerBuilder {
	sb.core.host = host

	return sb
}

func (sb *signerBuilder) Get(uri string) *signerBuilder {
	sb.core.method = methodGet
	sb.core.uri = uri

	return sb
}

func (sb *signerBuilder) Post(uri string) *signerBuilder {
	sb.core.method = methodPost
	sb.core.uri = uri

	return sb
}

func (sb *signerBuilder) Header(key string, value string) *signerBuilder {
	sb.core.headers[key] = value

	return sb
}

func (sb *signerBuilder) Headers(headers map[string]string) *signerBuilder {
	sb.core.headers = headers

	return sb
}

func (sb *signerBuilder) Payload(payload []byte) *signerBuilder {
	sb.core.payload = payload

	return sb
}

func (sb *signerBuilder) Product(product string) *signerBuilder {
	sb.params.product = product

	return sb
}

func (sb *signerBuilder) Service(service string) *signerBuilder {
	sb.params.service = service

	return sb
}

func (sb *signerBuilder) Zinan() *zinanBuilder {
	return newZinanBuilder(sb.params, sb.core, sb.getter)
}

func (sb *signerBuilder) Basic() *basicBuilder {
	return newBasicBuilder(sb.params, sb.core, sb.getter)
}

package songci

import (
	"strings"
)

type makerBuilder struct {
	params *params
	core   *coreParams
	getter getter
}

func newMakerBuilder(params *params, getter getter) *makerBuilder {
	return &makerBuilder{
		params: params,
		core:   newCoreParams(),
		getter: getter,
	}
}

func (mb *makerBuilder) Get(uri string) *makerBuilder {
	mb.core.method = methodGet
	values := strings.Split(uri, interrogation)
	mb.core.url = values[0]
	if 2 == len(values) {
		mb.core.query = values[1]
	}

	return mb
}

func (mb *makerBuilder) Post(uri string) *makerBuilder {
	mb.core.method = methodPost
	mb.core.url = uri

	return mb
}

func (mb *makerBuilder) Header(key string, value string) *makerBuilder {
	mb.core.headers[key] = value

	return mb
}

func (mb *makerBuilder) Headers(headers map[string]string) *makerBuilder {
	mb.core.headers = headers

	return mb
}

func (mb *makerBuilder) Payload(payload []byte) *makerBuilder {
	mb.core.payload = payload

	return mb
}

func (mb *makerBuilder) Product(product string) *makerBuilder {
	mb.params.product = product

	return mb
}

func (mb *makerBuilder) Service(service string) *makerBuilder {
	mb.params.service = service

	return mb
}

func (mb *makerBuilder) Zinan() *zinanBuilder {
	return newZinanBuilder(mb.params, mb.core, mb.getter)
}

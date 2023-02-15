package songci

type zinanSignerBuilder struct {
	signed []string
}

func newZinanSignerBuilder() *zinanSignerBuilder {
	return &zinanSignerBuilder{
		signed: make([]string, 0),
	}
}

func (zsb *zinanSignerBuilder) SignedHeader(header string) *zinanSignerBuilder {
	zsb.signed = append(zsb.signed, header)

	return zsb
}

func (zsb *zinanSignerBuilder) SignedHeaders(headers ...string) *zinanSignerBuilder {
	zsb.signed = headers

	return zsb
}

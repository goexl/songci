package songci

type maker struct {
	authorizer authorizer
}

func newMaker(authorizer authorizer) *maker {
	return &maker{
		authorizer: authorizer,
	}
}

func (m *maker) Make() (token string, codes []uint8) {
	return m.authorizer.token()
}

func (m *maker) Scheme() string {
	return m.authorizer.scheme()
}

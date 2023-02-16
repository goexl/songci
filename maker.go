package songci

import (
	"fmt"
)

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

func (m *maker) Auth() (auth string, codes []uint8) {
	if token, tc := m.Make(); nil != tc {
		codes = tc
	} else {
		auth = fmt.Sprintf("%s %s", m.Scheme(), token)
	}

	return
}

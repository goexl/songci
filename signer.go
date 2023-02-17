package songci

import (
	"fmt"
)

type signer struct {
	core       *coreParams
	authorizer authorizer
	auth       string
}

func newSigner(core *coreParams, authorizer authorizer) *signer {
	return &signer{
		core:       core,
		authorizer: authorizer,
	}
}

func (s *signer) Make() (token string, codes []uint8) {
	return s.authorizer.token()
}

func (s *signer) Scheme() string {
	return s.authorizer.scheme()
}

func (s *signer) Auth() (auth string, codes []uint8) {
	if "" != s.auth {
		auth = s.auth
	} else if token, tc := s.Make(); nil != tc {
		codes = tc
	} else {
		auth = fmt.Sprintf("%s %s", s.Scheme(), token)
		s.auth = auth
	}

	return
}

func (s *signer) Code() *codeBuilder {
	return newCodeBuilder(s.core, s.Auth)
}

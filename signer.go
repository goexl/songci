package songci

import (
	"fmt"
)

type signer struct {
	core          *coreParams
	authorizer    authorizer
	authorization string
}

func newSigner(core *coreParams, authorizer authorizer) *signer {
	return &signer{
		core:       core,
		authorizer: authorizer,
	}
}

func (s *signer) Credential() (string, []uint8) {
	return s.authorizer.credential()
}

func (s *signer) Scheme() string {
	return s.authorizer.scheme()
}

func (s *signer) Authorization() (authorization string, codes []uint8) {
	if "" != s.authorization {
		authorization = s.authorization
	} else if token, tc := s.Credential(); nil != tc {
		codes = tc
	} else {
		authorization = fmt.Sprintf("%s %s", s.Scheme(), token)
		s.authorization = authorization
	}

	return
}

func (s *signer) Code() *codeBuilder {
	return newCodeBuilder(s.core, s.Authorization)
}

package songci

import (
	"strings"

	"github.com/goexl/gox"
)

type verifier struct {
	params     *params
	self       *verifierParams
	token      string
	authorizer authorizer
}

func newVerifier(token string, params *params, self *verifierParams) *verifier {
	return &verifier{
		token:  token,
		params: params,
		self:   self,
	}
}

func (v *verifier) Verify() (codes []uint8) {
	values := strings.Split(v.token, space)
	switch values[0] {
	case v.params.zinan.name:
		v.authorizer = newZinan(v.params, newZinanParams(v.self))
	default:
		codes = append(codes, codeNotImplement)
	}
	if nil != codes {
		return
	}

	if uc := v.authorizer.unzip(v.token); nil != uc {
		codes = uc
	} else if signature, sc := v.authorizer.sign(); nil != sc {
		codes = sc
	} else {
		codes = gox.If(signature == v.authorizer.signature(), append(codes, codeSignatureError))
	}

	return
}

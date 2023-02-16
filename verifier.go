package songci

import (
	"strings"

	"github.com/goexl/gox"
)

type verifier struct {
	params     *params
	self       *verifierParams
	authorizer authorizer
}

func newVerifier(params *params, self *verifierParams) *verifier {
	return &verifier{
		params: params,
		self:   self,
	}
}

func (v *verifier) Verify(auth string) (codes []uint8) {
	values := strings.Split(auth, space)
	switch values[0] {
	case v.params.zinan.scheme:
		v.authorizer = newZinan(v.params, newZinanParams(v.self))
	default:
		codes = append(codes, codeNotImplement)
	}
	if nil != codes {
		return
	}

	if uc := v.authorizer.unzip(auth); nil != uc {
		codes = uc
	} else if signature, sc := v.authorizer.sign(); nil != sc {
		codes = sc
	} else {
		codes = gox.If(signature != v.authorizer.signature(), append(codes, codeSignatureError))
	}

	return
}

package songci

import (
	"github.com/goexl/gox"
)

type verifier struct {
	params     *params
	token      string
	authorizer authorizer
}

func newVerifier(params *params, authorizer authorizer) *verifier {
	return &verifier{
		params:     params,
		authorizer: authorizer,
	}
}

func (v *verifier) Verify() (codes []uint8) {
	if uc := v.authorizer.unzip(v.token); nil != uc {
		codes = uc
	} else if signature, sc := v.authorizer.sign(); nil != sc {
		codes = sc
	} else {
		codes = gox.If(signature == v.authorizer.signature(), append(codes, codeSignatureError))
	}

	return
}

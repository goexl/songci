package songci

import (
	"strings"

	"github.com/goexl/gox"
)

type verifier struct {
	params     *params
	core       *coreParams
	getter     getter
	authorizer authorizer
}

func newVerifier(params *params, core *coreParams, getter getter) *verifier {
	return &verifier{
		params: params,
		core:   core,
		getter: getter,
	}
}

func (v *verifier) Verify(auth string) (product string, service string, codes []uint8) {
	values := strings.Split(auth, space)
	switch values[0] {
	case v.params.zinan.scheme:
		v.authorizer = newZinan(v.params, v.core, newZinanParams(v.params, v.core), v.getter)
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
		product = v.params.product
		service = v.params.service
		codes = gox.If(signature != v.authorizer.signature(), append(codes, codeSignatureError))
	}

	return
}

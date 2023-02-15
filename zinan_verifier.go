package songci

import (
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

var _ Verifier = (*zinanVerifier)(nil)

type zinanVerifier struct {
	params *zinanParams
}

func newZinanVerifier(params *zinanParams) *zinanVerifier {
	return &zinanVerifier{
		params: params,
	}
}

func (zv *zinanVerifier) Verify() (verified bool, err error) {
	// 第一步，验证参数
	if codes, verified := zv.params.validate(); !verified {
		err = exc.NewField("参数不合法", field.New("codes", codes))
	}
	if nil != err {
		return
	}

	// 第二步，

	return
}

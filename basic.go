package songci

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var _ authorizer = (*basic)(nil)

type basic struct {
	params *params
	core   *coreParams
	self   *basicParams
	getter getter

	signature string
}

func newBasic(params *params, core *coreParams, self *basicParams, getter getter) *basic {
	return &basic{
		params: params,
		core:   core,
		self:   self,
		getter: getter,
	}
}

func (b *basic) scheme() string {
	return b.params.basic.scheme
}

func (b *basic) resolve(authorization string) (codes []uint8) {
	values := strings.Split(authorization, comma)
	b.params.id = strings.TrimSpace(strings.TrimPrefix(values[0], b.params.basic.scheme))
	if signature, de := base64.StdEncoding.DecodeString(values[1]); nil != de {
		codes = append(codes, codeSignatureError)
	} else {
		b.signature = string(signature)
	}

	return
}

func (b *basic) sign() (signature string, codes []uint8) {
	sign := fmt.Sprintf("%s:%s", b.params.id, b.self.password)
	signature = base64.StdEncoding.EncodeToString([]byte(sign))

	return
}

func (b *basic) token() (token string, codes []uint8) {
	if signature, _codes := b.sign(); nil != _codes {
		codes = _codes
	} else {
		token = signature
	}

	return
}

func (b *basic) check(signature string) bool {
	return signature == b.signature
}

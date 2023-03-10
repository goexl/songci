package songci

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/goexl/cryptor"
	"github.com/goexl/gox"
)

var _ authorizer = (*zinan)(nil)

type zinan struct {
	params *params
	core   *coreParams
	self   *zinanParams
	getter getter

	scope     string
	signed    string
	signature string
	secret    string
}

func newZinan(params *params, core *coreParams, self *zinanParams, getter getter) *zinan {
	return &zinan{
		params: params,
		core:   core,
		self:   self,
		getter: getter,
	}
}

func (z *zinan) scheme() string {
	return z.params.zinan.scheme
}

func (z *zinan) resolve(authorization string) (codes codes) {
	values := strings.Split(authorization, comma)
	if zinanAuthorizationSize != len(values) {
		codes.Add(1)
	} else {
		codes = z.resolveValues(values)
	}

	return
}

func (z *zinan) resolveValues(values []string) (codes codes) {
	z.params.id = strings.TrimSpace(strings.TrimPrefix(values[0], z.params.zinan.scheme))
	if _codes := z.self.resolveScope(values[1]); nil != _codes {
		codes = _codes
	} else {
		timestamp := time.Unix(z.self.timestamp, 0)
		checked := time.Since(timestamp).Abs() <= z.params.timeout
		codes = gox.If(!checked, append(codes, codeTimeout))
		z.self.resolveSigned(values[2])
		z.signature = values[3]
	}

	return
}

func (z *zinan) sign() (signature string, codes codes) {
	if vc := z.self.validate(); nil != vc {
		codes = vc
	} else if secret, err := z.getter.Get(z.params.zinan.scheme, z.params.id); nil != err {
		codes = append(codes, codeGetCredentialError)
	} else {
		z.secret = secret
	}
	if nil != codes {
		return
	}

	timestamp := fmt.Sprintf("%d", z.self.timestamp)
	z.signed = z.self.signedHeaders()
	z.scope = z.self.scope()

	req := new(strings.Builder)
	// 写入方法
	req.WriteString(z.core.method)
	req.WriteString(enter)
	values := strings.Split(z.core.uri, interrogation)
	// 写入请求地址
	req.WriteString(values[0])
	req.WriteString(enter)
	if 2 == len(values) {
		// 写入查询参数
		req.WriteString(url.QueryEscape(values[1]))
		req.WriteString(enter)
	}
	// 写入头
	req.WriteString(z.self.processedHeaders())
	req.WriteString(enter)
	// 写入签名头
	req.WriteString(z.signed)
	req.WriteString(enter)
	// 写入有效荷载
	req.WriteString(cryptor.New(z.self.payload).Sha().Build().Hex())

	sign := new(strings.Builder)
	// 写入算法名
	sign.WriteString(z.params.zinan.scheme)
	sign.WriteString(enter)
	// 写入时间戳
	sign.WriteString(timestamp)
	sign.WriteString(enter)
	// 写入作用域
	sign.WriteString(z.scope)
	sign.WriteString(enter)
	// 写入请求
	sign.WriteString(cryptor.New(req.String()).Sha().Build().Hex())

	_timestamp := cryptor.New(timestamp).Hmac(z.self.secret(z.secret)).Build().String()
	_service := cryptor.New(z.params.service).Hmac(_timestamp).Build().String()
	signing := cryptor.New(z.self.request()).Hmac(_service).Build().String()
	signature = cryptor.New(sign.String()).Hmac(signing).Build().Hex()

	return
}

func (z *zinan) token() (token string, codes codes) {
	if signature, _codes := z.sign(); nil != _codes {
		codes = _codes
	} else {
		sb := new(strings.Builder)
		// 写入应用编号
		sb.WriteString(z.params.id)
		sb.WriteString(comma)
		// 写入作用域
		sb.WriteString(z.scope)
		sb.WriteString(comma)
		// 写入签名头
		sb.WriteString(z.signed)
		sb.WriteString(comma)
		// 写入签名值
		sb.WriteString(signature)
		token = sb.String()
	}

	return
}

func (z *zinan) check(signature string) bool {
	return signature == z.signature
}

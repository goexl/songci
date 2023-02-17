package songci

import (
	"fmt"
	"math"
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

	scope          string
	signedHeaders  string
	finalSignature string
	credential     string
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

func (z *zinan) unzip(auth string) (codes []uint8) {
	values := strings.Split(auth, comma)
	z.params.id = strings.TrimSpace(strings.TrimPrefix(values[0], z.params.zinan.scheme))
	if _codes := z.self.unzipScope(values[1]); nil != _codes {
		codes = _codes
	} else {
		timestamp := time.Unix(z.self.timestamp, 0)
		checked := time.Duration(math.Abs(float64(time.Now().Sub(timestamp)))) <= z.params.timeout
		codes = gox.If(!checked, append(codes, codeTimeout))
		z.self.unzipSigned(values[2])
		z.finalSignature = values[3]
	}

	return
}

func (z *zinan) sign() (signature string, codes []uint8) {
	if vc := z.self.validate(); nil != vc {
		codes = vc
	} else if credential, err := z.getter.Get(z.params.zinan.scheme, z.params.id); nil != err {
		codes = append(codes, codeGetCredentialError)
	} else {
		z.credential = credential
	}
	if nil != codes {
		return
	}

	timestamp := fmt.Sprintf("%d", z.self.timestamp)
	z.signedHeaders = z.self.signedHeaders()
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
		req.WriteString(values[1])
		req.WriteString(enter)
	}
	// 写入头
	req.WriteString(z.self.processedHeaders())
	req.WriteString(enter)
	// 写入签名头
	req.WriteString(z.signedHeaders)
	req.WriteString(enter)
	// 写入有效荷载
	req.WriteString(cryptor.New(z.self.payload).Sha256().Hex())

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
	sign.WriteString(cryptor.New(req.String()).Sha256().Hex())

	secret := cryptor.New(timestamp).Hmac(z.self.secret(z.credential)).String()
	svc := cryptor.New(z.params.service).Hmac(secret).String()
	signing := cryptor.New(z.self.request()).Hmac(svc).String()
	signature = cryptor.New(sign.String()).Hmac(signing).Hex()

	return
}

func (z *zinan) token() (token string, codes []uint8) {
	sb := new(strings.Builder)
	if signature, _codes := z.sign(); nil != _codes {
		codes = _codes
	} else {
		// 写入应用编号
		sb.WriteString(z.params.id)
		sb.WriteString(comma)
		// 写入作用域
		sb.WriteString(z.scope)
		sb.WriteString(comma)
		// 写入签名头
		sb.WriteString(z.signedHeaders)
		sb.WriteString(comma)
		// 写入签名值
		sb.WriteString(signature)
		token = sb.String()
	}

	return
}

func (z *zinan) signature() string {
	return z.finalSignature
}

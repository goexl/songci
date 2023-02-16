package songci

import (
	"fmt"
	"strings"
	"time"

	"github.com/goexl/cryptor"
)

var _ authorizer = (*zinan)(nil)

type zinan struct {
	params *params
	self   *zinanParams

	scope  string
	signed string
}

func (z *zinan) sign() (signature string, err error) {
	timestamp := time.Now().Unix()
	z.signed = z.self.signed()
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	z.scope = z.self.scope(date)

	request := new(strings.Builder)
	// 写入方法
	request.WriteString(z.self.method)
	request.WriteString(enter)
	// 写入请求地址
	request.WriteString(z.self.uri)
	request.WriteString(enter)
	// 写入查询参数
	request.WriteString(z.self.query)
	request.WriteString(enter)
	// 写入头
	request.WriteString(z.self.headers())
	request.WriteString(enter)
	// 写入签名头
	request.WriteString(z.signed)
	request.WriteString(enter)
	// 写入有效荷载
	request.WriteString(cryptor.New(z.self.payload).Sha256().Hex())

	sign := new(strings.Builder)
	// 写入算法名
	sign.WriteString(z.self.algorithm)
	sign.WriteString(enter)
	// 写入时间戳
	sign.WriteString(fmt.Sprintf("%d", timestamp))
	sign.WriteString(enter)
	// 写入作用域
	sign.WriteString(z.scope)
	sign.WriteString(enter)
	// 写入请求
	sign.WriteString(cryptor.New(request.String()).Sha256().Hex())

	secret := cryptor.New(date).Hmac(z.self.secret()).String()
	service := cryptor.New(z.self.service).Hmac(secret).String()
	signing := cryptor.New(z.self.request()).Hmac(service).String()
	signature = cryptor.New(sign.String()).Hmac(signing).Hex()

	return
}

func (z *zinan) authorize() (authorize string, err error) {
	sb := new(strings.Builder)
	if signature, se := z.sign(); nil != se {
		err = se
	} else {
		// 写入应用编号
		sb.WriteString(z.self.id)
		sb.WriteString(comma)
		// 写入作用域
		sb.WriteString(z.scope)
		sb.WriteString(comma)
		// 写入签名头
		sb.WriteString(z.signed)
		sb.WriteString(comma)
		// 写入签名值
		sb.WriteString(signature)
		authorize = sb.String()
	}

	return
}

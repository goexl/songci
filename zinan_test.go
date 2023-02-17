package songci_test

import (
	"testing"

	"github.com/goexl/gox/rand"
	"github.com/goexl/songci"
)

const zinanLen = 26

func TestZinan(t *testing.T) {
	secret := songci.NewSecret(rand.New().String().Length(zinanLen).Generate())
	headers := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"Host":         host,
	}
	zinan := songci.New().Build().Singer(id, secret).Host(host).Get(getUri).Headers(headers).Zinan().Build()
	verifier := songci.New().Build().Verifier(secret).Get().Uri(getUri).Headers(headers).Build()
	if auth, mc := zinan.Authorization(); nil != mc {
		t.Errorf("签名测试未通过，密钥：%s，错误：%v", secret, mc)
	} else if _, _, vc := verifier.Verify(auth); nil != vc {
		t.Errorf("验签测试未通过，密钥：%s，错误：%v", secret, vc)
	}
}

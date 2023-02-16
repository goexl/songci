package songci_test

import (
	"testing"

	"github.com/goexl/gox/rand"
	"github.com/goexl/songci"
)

const zinanLen = 26

func TestZinan(t *testing.T) {
	credential := rand.New().String().Length(zinanLen).Generate()
	zinan := songci.New().Build().Maker(credential).Zinan().Get(uri)
	verifier := songci.New().Build().Verifier(credential).Get().Uri(uri)
	if token, mc := zinan.Build().Make(); nil != mc {
		t.Errorf("签名测试未通过，密钥：%s，错误：%v", credential, mc)
	} else if vc := verifier.Build().Verify(token); nil != vc {
		t.Errorf("验签测试未通过，密钥：%s，错误：%v", credential, mc)
	}
}

package songci

import (
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

var (
	_ getter = (*mapGetter)(nil)
	_        = NewMapSecret
)

type mapGetter struct {
	secrets map[string]string
}

// NewMapSecret 基于映射的凭据获取
func NewMapSecret(secrets map[string]string) *mapGetter {
	return &mapGetter{
		secrets: secrets,
	}
}

func (mg *mapGetter) Get(_ string, id string) (secret string, err error) {
	if _secret, ok := mg.secrets[id]; ok {
		secret = _secret
	} else {
		err = exc.NewField("找不到对应的凭据", field.New("id", id))
	}

	return
}

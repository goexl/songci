package songci

import (
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

var (
	_ getter = (*mapGetter)(nil)
	_        = NewMapCredential
)

type mapGetter struct {
	credentials map[string]string
}

// NewMapCredential 基于映射的凭据获取
func NewMapCredential(credentials map[string]string) *mapGetter {
	return &mapGetter{
		credentials: credentials,
	}
}

func (sg *mapGetter) Get(_ string, id string) (credential string, err error) {
	if _credential, ok := sg.credentials[id]; ok {
		credential = _credential
	} else {
		err = exc.NewField("找不到对应的凭据", field.New("id", id))
	}

	return
}

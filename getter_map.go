package songci

import (
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

var (
	_ getter = (*mapGetter)(nil)
	_        = NewMapGetter
)

type mapGetter struct {
	credentials map[string]string
}

// NewMapGetter 基于映射的凭据获取
func NewMapGetter(credentials map[string]string) *mapGetter {
	return &mapGetter{
		credentials: credentials,
	}
}

func (sg *mapGetter) Get(id string) (credential string, err error) {
	if _credential, ok := sg.credentials[id]; ok {
		credential = _credential
	} else {
		err = exc.NewField("找不到对应的凭据", field.New("id", id))
	}

	return
}

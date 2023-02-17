package songci

import (
	"fmt"
	"sort"
	"strings"

	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

type coder struct {
	core    *coreParams
	params  *codeParams
	authFun authFun
}

func newCoder(core *coreParams, params *codeParams, authFun authFun) *coder {
	return &coder{
		core:    core,
		params:  params,
		authFun: authFun,
	}
}

func (c *coder) String() (code string, err error) {
	switch c.params.typ {
	case codeTypeCurl:
		code, err = c.curl()
	}

	return
}

func (c *coder) curl() (command string, err error) {
	cmd := new(commands)
	cmd.append(curl)
	cmd.append(c.bashEscape(c.core.method))
	if nil != c.core.payload {
		cmd.append(curlData)
		cmd.append(c.bashEscape(string(c.core.payload)))
	}

	// 复制一份新的数据，防止改变老数据
	_headers := make(headers, len(c.core.headers)+1)
	keys := make([]string, 0, len(c.core.headers))
	for key, value := range c.core.headers {
		keys = append(keys, key)
		_headers[key] = value
	}
	if auth, ac := c.authFun(); nil != ac {
		err = exc.NewField("生成授权出错", field.New("codes", ac))
	} else {
		_headers[authorization] = auth
		keys = append(keys, authorization)
		sort.Strings(keys)
	}
	if nil != err {
		return
	}

	for _, key := range keys {
		cmd.append(curlHeader)
		cmd.append(c.bashEscape(fmt.Sprintf("%s: %s", key, _headers[key])))
	}
	cmd.append(c.bashEscape(fmt.Sprintf("%s://%s%s", c.params.scheme, c.core.host, c.core.uri)))
	command = cmd.String()

	return
}

func (c *coder) bashEscape(from string) string {
	return `'` + strings.Replace(from, `'`, `'\''`, -1) + `'`
}

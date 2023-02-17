package songci

import (
	"fmt"
	"sort"
	"strings"
)

type maker struct {
	core       *coreParams
	authorizer authorizer
}

func newMaker(core *coreParams, authorizer authorizer) *maker {
	return &maker{
		core:       core,
		authorizer: authorizer,
	}
}

func (m *maker) Make() (token string, codes []uint8) {
	return m.authorizer.token()
}

func (m *maker) Scheme() string {
	return m.authorizer.scheme()
}

func (m *maker) Auth() (auth string, codes []uint8) {
	if token, tc := m.Make(); nil != tc {
		codes = tc
	} else {
		auth = fmt.Sprintf("%s %s", m.Scheme(), token)
	}

	return
}

func (m *maker) Curl() string {
	command := new(strings.Builder)
	command.WriteString(curl)
	command.WriteString(m.bashEscape(m.core.method))
	command.WriteString(curlData)
	command.WriteString(m.bashEscape(string(m.core.payload)))

	keys := make([]string, 0, len(m.core.headers))
	for key := range m.core.headers {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		command.WriteString(curlHeader)
		command.WriteString(m.bashEscape(fmt.Sprintf("%s: %s", key, m.core.headers[key])))
	}
	command.WriteString(m.bashEscape(fmt.Sprintf("%s%s", m.core.host, m.core.url)))

	return command.String()
}

func (m *maker) bashEscape(from string) string {
	return `'` + strings.Replace(from, "'", `'\''`, -1) + `'`
}

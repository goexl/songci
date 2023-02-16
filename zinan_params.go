package songci

import (
	"fmt"
	"sort"
	"strings"
)

type zinanParams struct {
	algorithm string
	id        string
	_secret   string
	service   string
	product   string
	version   int
	method    string
	uri       string
	query     string
	original  headers
	processed headers
	_signed   []string
	payload   []byte
}

func newZinanParams() *zinanParams {
	return &zinanParams{
		algorithm: zinanName,
		method:    methodPost,
		uri:       rootPath,
		original:  make(map[string]string),
	}
}

func (zp *zinanParams) secret() (final string) {
	sb := new(strings.Builder)
	if "" != zp.product {
		sb.WriteString(strings.ToUpper(zp.product))
	}
	if 0 != zp.version {
		sb.WriteString(fmt.Sprintf("%d", zp.version))
	}
	sb.WriteString(zp._secret)

	return
}

func (zp *zinanParams) request() (final string) {
	sb := new(strings.Builder)
	if "" != zp.product {
		sb.WriteString(strings.ToLower(zp.product))
	}
	if 0 != zp.version {
		sb.WriteString(fmt.Sprintf("%d", zp.version))
	}
	sb.WriteString(underline)
	sb.WriteString(request)

	return
}

func (zp *zinanParams) scope(date string) string {
	return fmt.Sprintf("%s/%s/%s", date, zp.service, zp.request())
}

func (zp *zinanParams) headers() (headers string) {
	keys := make([]string, 0, len(zp.processed))
	for k := range zp.processed {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sb := new(strings.Builder)
	for _, key := range keys {
		sb.WriteString(key)
		sb.WriteString(equal)
		sb.WriteString(zp.processed[key])
		sb.WriteString(semicolon)
	}

	return sb.String()[:sb.Len()-1]
}

func (zp *zinanParams) signed() (signed string) {
	sort.Strings(zp._signed)
	signed = strings.Join(zp._signed, semicolon)

	return
}

func (zp *zinanParams) validate() (codes []int, validate bool) {
	hasContentType := false
	hasHost := false
	zp.processed = make(headers, len(zp.original))
	for key, value := range zp.original {
		newKey := strings.ToLower(key)
		if contentType == newKey {
			hasContentType = true
		}
		if host == newKey {
			hasHost = true
		}
		zp.processed[newKey] = value
	}

	validate = hasHost && hasContentType
	if !hasContentType {
		codes = append(codes, codeNoContentTypeHeader)
	}
	if !hasHost {
		codes = append(codes, codeNoHostHeader)
	}

	return
}

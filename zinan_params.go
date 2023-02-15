package songci

import (
	"strings"
)

type zinanParams struct {
	method  string
	uri     string
	query   string
	headers headers
	payload []byte
}

func newZinanParams() *zinanParams {
	return &zinanParams{
		method:  methodPost,
		uri:     rootPath,
		headers: make(map[string]string),
	}
}

func (zp *zinanParams) validate() (codes []int, validate bool) {
	hasContentType := false
	hasHost := false
	for key, value := range zp.headers {
		newKey := strings.ToLower(key)
		value = strings.ToLower(value)
		if contentType == newKey {
			hasContentType = true
		}
		if host == key {
			hasHost = true
		}
		delete(zp.headers, key)
		zp.headers[newKey] = value
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

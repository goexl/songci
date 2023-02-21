package songci

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type zinanParams struct {
	core      *coreParams
	params    *params
	version   string
	signed    []string
	payload   []byte
	timestamp int64
}

func newZinanParams(params *params, core *coreParams) *zinanParams {
	return &zinanParams{
		core:      core,
		params:    params,
		version:   "1",
		signed:    []string{contentType, userAgent},
		timestamp: time.Now().Unix(),
	}
}

func (zp *zinanParams) secret(credential string) string {
	sb := new(strings.Builder)
	sb.WriteString(zp.params.product)
	sb.WriteString(zp.version)
	sb.WriteString(credential)

	return sb.String()
}

func (zp *zinanParams) request() string {
	sb := new(strings.Builder)
	sb.WriteString(zp.params.service)
	sb.WriteString(underline)
	sb.WriteString(zp.version)
	sb.WriteString(underline)
	sb.WriteString(request)

	return sb.String()
}

func (zp *zinanParams) resolveRequest(request string) {
	values := strings.Split(request, underline)
	zp.params.service = values[0]
	zp.version = values[1]
}

func (zp *zinanParams) scope() string {
	sb := new(strings.Builder)
	sb.WriteString(fmt.Sprintf("%d", zp.timestamp))
	sb.WriteString(slash)
	sb.WriteString(zp.params.product)
	sb.WriteString(slash)
	sb.WriteString(zp.request())

	return sb.String()
}

func (zp *zinanParams) resolveScope(scope string) (codes codes) {
	values := strings.Split(scope, slash)
	if number, pe := strconv.ParseInt(values[0], 10, 64); nil != pe {
		codes = append(codes, codeTimestampFormatError)
	} else {
		zp.timestamp = number
		zp.params.product = values[1]
		zp.resolveRequest(values[2])
	}

	return
}

func (zp *zinanParams) processedHeaders() (headers string) {
	sort.Strings(zp.signed)
	sb := new(strings.Builder)
	for _, key := range zp.signed {
		sb.WriteString(key)
		sb.WriteString(equal)
		sb.WriteString(zp.core.headers[key])
		sb.WriteString(semicolon)
	}

	return sb.String()[:sb.Len()-1]
}

func (zp *zinanParams) signedHeaders() (signed string) {
	sort.Strings(zp.signed)
	signed = strings.Join(zp.signed, semicolon)

	return
}

func (zp *zinanParams) resolveSigned(signed string) {
	zp.signed = strings.Split(signed, semicolon)
}

func (zp *zinanParams) validate() (codes codes) {
	hasContentType := false
	hasUserAgent := false
	for key, value := range zp.core.headers {
		zp.core.headers[strings.ToLower(key)] = value
		delete(zp.core.headers, key)

		if contentType == key {
			hasContentType = true
		}
		if userAgent == key {
			hasUserAgent = true
		}
	}

	if !hasContentType {
		codes.Add(codeNoContentTypeHeader)
	}
	if !hasUserAgent {
		codes.Add(codeNoUserAgentHeader)
	}

	return
}

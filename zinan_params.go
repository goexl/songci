package songci

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

func (zp *zinanParams) validate() bool {
	hasContentType := false
	hasHost := false
	for key := range zp.headers {
		if contentType == key {
			hasContentType = true
		}
		if host == key {
			hasHost = true
		}
	}

	return hasContentType && hasHost
}
